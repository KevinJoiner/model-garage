package convert

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

type FunctionInfo struct {
	Comments string
	Body     []byte
}

func createConvertFuncs(tmplData *schema.TemplateData, outputDir string, copyComments bool, convertFunc []conversionData, existingFuncs map[string]FunctionInfo) error {
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

// getConversionFunctions returns the signals that need conversion functions and test functions.
func getConversionFunctions(signals []*schema.SignalInfo) ([]conversionData, []conversionData) {
	var convertFunc []conversionData
	var convertTestFunc []conversionData
	for _, signal := range signals {
		if len(signal.Conversions) == 0 {
			continue
		}
		for i := range signal.Conversions {
			funcName := convertName(signal) + strconv.Itoa(i)
			convData := conversionData{Signal: signal, convIdx: i, FuncName: funcName}
			convertFunc = append(convertFunc, convData)

			funcName = convertTestName(signal) + strconv.Itoa(i)
			convData = conversionData{Signal: signal, convIdx: i, FuncName: funcName}
			convertTestFunc = append(convertTestFunc, convData)

		}
	}
	return convertFunc, convertTestFunc
}

// createConvertTestFunc creates test functions for the conversion functions if they do not exist.
func createConvertTestFunc(tmplData *schema.TemplateData, outputDir string, copyComments bool, convertTestFunc []conversionData, existingFuncs map[string]FunctionInfo) error {
	convertTestFuncTemplate, err := createConvertTestFuncTemplate(tmplData.PackageName)
	if err != nil {
		return err
	}

	if len(convertTestFunc) == 0 {
		return nil
	}

	convertTestFuncFileName := fmt.Sprintf(convertTestFuncFileNameFormat, strings.ToLower(tmplData.ModelName))
	filePath := filepath.Join(outputDir, convertTestFuncFileName)
	packageName := tmplData.PackageName + "_test"
	err = writeConvertFuncs(convertTestFunc, existingFuncs, convertTestFuncTemplate, filePath, packageName, copyComments)
	if err != nil {
		return err
	}
	return nil
}

func createConvertFuncTemplate() (*template.Template, error) {
	tmpl, err := template.New("convertFuncTemplate").Parse(convertFuncTemplateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}

func createConvertTestFuncTemplate(packageNameToTest string) (*template.Template, error) {
	tmpl, err := template.New("convertTestFuncTemplate").Funcs(template.FuncMap{
		"trimSuffix": strings.TrimSuffix,
	}).Parse(convertTestsFuncTemplateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}

func getDeclaredFunctions(outputPath string) (map[string]FunctionInfo, error) {
	fset := token.NewFileSet()
	declaredFunctions := make(map[string]FunctionInfo)

	list, err := os.ReadDir(outputPath)
	if err != nil {
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			continue
		}
		filename := filepath.Join(outputPath, d.Name())
		err = addFileDeclarations(fset, filename, declaredFunctions)
		if err != nil {
			return nil, err
		}
	}

	return declaredFunctions, nil
}

func addFileDeclarations(fset *token.FileSet, filePath string, declaredFunctions map[string]FunctionInfo) error {
	src, err := parser.ParseFile(fset, filePath, nil, parser.ParseComments)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
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
			Node:     fn,
			Comments: src.Comments,
		})
		if err != nil {
			return fmt.Errorf("error formating function: %w", err)
		}

		declaredFunctions[fn.Name.Name] = FunctionInfo{
			Comments: strings.Join(docComments, "\n"),
			Body:     buf.Bytes(),
		}
	}
	return nil
}

func convertName(signal *schema.SignalInfo) string {
	return "To" + signal.GOName
}

func convertTestName(signal *schema.SignalInfo) string {
	return "Test" + convertName(signal)
}

func writeConvertFuncs(convertFunc []conversionData, existingFuncs map[string]FunctionInfo, tmpl *template.Template, outputPath string, packageName string, copyComments bool) error {
	var convertBuff bytes.Buffer
	convertBuff.WriteString(fmt.Sprintf(header, packageName))
	// Add or update existing functions
	slices.SortFunc(convertFunc, func(a, b conversionData) int {
		return cmp.Compare(a.FuncName, b.FuncName)
	})

	for _, convData := range convertFunc {
		funcName := convData.FuncName
		var docComment, body string
		if fnInfo, exists := existingFuncs[funcName]; exists {
			body = string(fnInfo.Body)
			if copyComments {
				docComment = fnInfo.Comments
			}
		}
		err := tmpl.Execute(&convertBuff, funcTmplData{
			PackageName: packageName,
			Signal:      convData.Signal,
			FuncName:    funcName,
			Conversion:  convData.Signal.Conversions[convData.convIdx],
			DocComment:  docComment,
			Body:        body,
		})
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
