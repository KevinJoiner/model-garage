package convert

import (
	"bytes"
	_ "embed"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/KevinJoiner/model-garage/internal/codegen"
)

//go:embed convert.tmpl
var convertTemplate string

//go:embed convertFunc.tmpl
var convertFuncTemplate string

const header = `package %s

// This file is automatically populated with conversion functions for each field of a vehicle struct.
// any conversion functions already defined in this package will not be generated.
// Code generated by model-garage.
`

// Generate creates a conversion functions for each field of a vehicle struct.
// as well as the entire vehicle struct.
func Generate(tmplData *codegen.TemplateData, outputDir string) (err error) {
	goTemplate, err := createGoTemplate()
	if err != nil {
		return err
	}
	convertFuncTemplate, err := createConvertFuncTemplate()
	if err != nil {
		return err
	}

	// execute the struct template
	var outBuf bytes.Buffer
	if err := goTemplate.Execute(&outBuf, &tmplData); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}

	goOutputPath := filepath.Join(outputDir, codegen.ConvertFileName)
	// format and write the go file.
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), goOutputPath)
	if err != nil {
		return fmt.Errorf("error formatting and writing to file: %w", err)
	}

	existingFuncs, err := getDeclaredFunctions(outputDir)
	if err != nil {
		return fmt.Errorf("error getting declared functions: %w", err)
	}

	// get a list of SignalInfos that need new convert functions
	var needsConvertFunc []*codegen.SignalInfo
	for _, signal := range tmplData.DataSignals {
		if signal.Conversion != nil && !existingFuncs[convertName(signal)] {
			needsConvertFunc = append(needsConvertFunc, signal)
		}
	}
	if len(needsConvertFunc) == 0 {
		return nil
	}

	// check if we need to create convertFunc file
	convertFuncPath := filepath.Join(outputDir, codegen.ConvertFuncFileName)

	err = ensureFuncFile(convertFuncPath, tmplData.PackageName)
	if err != nil {
		return err
	}

	convertData, err := os.ReadFile(convertFuncPath)
	if err != nil {
		return fmt.Errorf("error reading convertFunc file: %w", err)
	}
	convertBuff := bytes.NewBuffer(convertData)
	for _, signal := range needsConvertFunc {
		if err := convertFuncTemplate.Execute(convertBuff, signal); err != nil {
			return fmt.Errorf("error executing convertFunc template: %w", err)
		}
	}
	err = codegen.FormatAndWriteToFile(convertBuff.Bytes(), convertFuncPath)
	if err != nil {
		return fmt.Errorf("error formatting and writing to file: %w", err)
	}

	return nil
}

func createGoTemplate() (*template.Template, error) {
	tmpl, err := template.New("convertTemplate").Funcs(template.FuncMap{
		"convertName": convertName,
	}).Parse(convertTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}

func createConvertFuncTemplate() (*template.Template, error) {
	tmpl, err := template.New("convertFuncTemplate").Funcs(template.FuncMap{
		"convertName": convertName,
	}).Parse(convertFuncTemplate)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}

func getDeclaredFunctions(outputPath string) (map[string]bool, error) {
	fset := token.NewFileSet()
	declaredFunctions := map[string]bool{}

	list, err := os.ReadDir(outputPath)
	if err != nil {
		return nil, err
	}

	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			continue
		}
		filename := filepath.Join(outputPath, d.Name())
		src, err := parser.ParseFile(fset, filename, nil, parser.SkipObjectResolution|parser.ParseComments)
		if err != nil {
			return nil, fmt.Errorf("error parsing file: %w", err)
		}
		for _, decl := range src.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Recv != nil {
				continue
			}
			declaredFunctions[fn.Name.Name] = true
		}

	}

	return declaredFunctions, nil
}

func convertName(signal *codegen.SignalInfo) string {
	return "To" + signal.GOName
}

// ensureFuncFile checks if the convertFunc file exists and creates it if it does not.
// It also writes the package header to the file if it is created.
func ensureFuncFile(convertFuncPath string, packageName string) error {
	_, err := os.Stat(convertFuncPath)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("error checking for %s file: %w", codegen.ConvertFuncFileName, err)
	}
	// create the convertFunc file
	funcFile, err := os.Create(convertFuncPath)
	if err != nil {
		return fmt.Errorf("error creating convertFunc file: %w", err)
	}
	_, err = funcFile.WriteString(fmt.Sprintf(header, packageName))
	if err != nil {
		_ = funcFile.Close()
		return fmt.Errorf("error writing to convertFunc file: %w", err)
	}
	err = funcFile.Close()
	if err != nil {
		return fmt.Errorf("error closing convertFunc file: %w", err)
	}

	return nil
}
