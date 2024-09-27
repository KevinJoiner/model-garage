package convert

// This package provides functions for converting data between different formats.

import (
	"bytes"
	"cmp"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/printer"
	"go/token"
	"os"
	"path/filepath"
	"slices"
	"strconv"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/pkg/codegen"
	"github.com/DIMO-Network/model-garage/pkg/schema"
)

type functionInfo struct {
	Comments string
	Body     []byte
}

// createConvertFuncs generates conversion functions based on the provided template data.
// It writes the generated functions to the specified output directory.
// If copyComments is true, it will copy existing comments from the original functions.
// The existingFuncs map contains information about existing functions in the output directory.
func createConvertFuncs(tmplData *schema.TemplateData, outputDir string, copyComments bool, convertFunc []funcTmplData, existingFuncs map[string]functionInfo) error {
	convertFuncTemplate, err := createConvertFuncTemplate()
	if err != nil {
		return err
	}
	if len(convertFunc) == 0 {
		return nil
	}

	convertFuncFileName := fmt.Sprintf(convertFuncFileNameFormat, strings.ToLower(tmplData.ModelName))
	filePath := filepath.Join(outputDir, convertFuncFileName)
	err = writeConvertFuncs(convertFunc, existingFuncs, convertFuncTemplate, filePath, tmplData.PackageName, copyComments)
	if err != nil {
		return err
	}
	return nil
}

// getConversionFunctions returns the signals that need conversion functions.
func getConversionFunctions(signals []*schema.SignalInfo) []funcTmplData {
	var convertFunc []funcTmplData
	for _, signal := range signals {
		for i := range signal.Conversions {
			funcName := convertName(signal) + strconv.Itoa(i)
			convData := funcTmplData{
				Signal:     signal,
				Conversion: signal.Conversions[i],
				FuncName:   funcName,
			}

			convertFunc = append(convertFunc, convData)
		}
	}
	return convertFunc
}

// createConvertFuncTemplate creates a template for generating conversion functions.
func createConvertFuncTemplate() (*template.Template, error) {
	tmpl, err := template.New("convertFuncTemplate").Parse(convertFuncTemplateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}

// GetDeclaredFunctions returns a map of function names to their corresponding function information for a given directory.
// The function information includes comments and body.
func GetDeclaredFunctions(outputPath string) (map[string]functionInfo, error) {
	fset := token.NewFileSet()
	declaredFunctions := make(map[string]functionInfo)

	list, err := os.ReadDir(outputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			continue
		}
		filename := filepath.Join(outputPath, d.Name())
		fileDeclaredFunctions, err := getDeclaredFunctionsForFile(fset, filename)
		if err != nil {
			return nil, err
		}
		for k, v := range fileDeclaredFunctions {
			declaredFunctions[k] = v
		}
	}

	return declaredFunctions, nil
}

func getDeclaredFunctionsForFile(fset *token.FileSet, filePath string) (map[string]functionInfo, error) {
	declaredFunctions := make(map[string]functionInfo)
	src, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return nil, fmt.Errorf("error parsing file: %w", err)
	}
	for _, decl := range src.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv != nil {
			continue
		}

		var docComments []string
		if fn.Doc != nil {
			for _, comment := range fn.Doc.List {
				docComments = append(docComments, comment.Text)
			}
		}
		// set comments to nil to avoid printing them in the body
		fn.Doc = nil

		// Capture function body including comments
		var buf bytes.Buffer
		err = format.Node(&buf, fset, &printer.CommentedNode{
			Node:     fn.Body,
			Comments: src.Comments,
		})
		if err != nil {
			return nil, fmt.Errorf("error formating function: %w", err)
		}

		declaredFunctions[fn.Name.Name] = functionInfo{
			Comments: strings.Join(docComments, "\n"),
			Body:     buf.Bytes(),
		}
	}
	return declaredFunctions, nil
}

// convertName returns the conversion function name for a given signal.
func convertName(signal *schema.SignalInfo) string {
	return "To" + signal.GOName
}

// writeConvertFuncs writes the generated conversion functions to a file.
func writeConvertFuncs(convertFunc []funcTmplData, existingFuncs map[string]functionInfo, tmpl *template.Template, outputPath string, packageName string, copyComments bool) error {
	var convertBuff bytes.Buffer
	convertBuff.WriteString(fmt.Sprintf(header, packageName))
	slices.SortStableFunc(convertFunc, func(a, b funcTmplData) int {
		return cmp.Compare(a.FuncName, b.FuncName)
	})

	// Add or update existing functions
	for _, convData := range convertFunc {
		funcName := convData.FuncName
		if fnInfo, exists := existingFuncs[funcName]; exists {
			convData.Body = string(fnInfo.Body)
			if copyComments {
				convData.DocComment = fnInfo.Comments
			}
		}

		err := tmpl.Execute(&convertBuff, convData)
		if err != nil {
			return fmt.Errorf("error executing template for function %s: %w", funcName, err)
		}
	}

	err := codegen.FormatAndWriteToFile(convertBuff.Bytes(), outputPath)
	if err != nil {
		return fmt.Errorf("error formatting and writing to file: %w", err)
	}
	return nil
}
