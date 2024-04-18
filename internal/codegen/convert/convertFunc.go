package convert

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/internal/codegen"
)

func createConvertFuncs(tmplData *codegen.TemplateData, outputDir string, needsConvertFunc []ConversionData) error {
	convertFuncTemplate, err := createConvertFuncTemplate()
	if err != nil {
		return err
	}
	if len(needsConvertFunc) == 0 {
		return nil
	}

	convertFuncFileName := fmt.Sprintf(convertFuncFileNameFormat, strings.ToLower(tmplData.ModelName))
	filePath := filepath.Join(outputDir, convertFuncFileName)
	err = writeConvertFuncs(needsConvertFunc, convertFuncTemplate, filePath, tmplData.PackageName)
	if err != nil {
		return err
	}
	return nil
}

// getConversionFunctions returns the signals that need conversion functions and test functions.
func getConversionFunctions(signals []*codegen.SignalInfo, existingFuncs map[string]bool) ([]ConversionData, []ConversionData) {
	var needsConvertFunc []ConversionData
	var needsConvertTestFunc []ConversionData
	for _, signal := range signals {
		if len(signal.Conversions) == 0 {
			continue
		}
		for i, _ := range signal.Conversions {
			funcName := convertName(signal) + strconv.Itoa(i)
			if !existingFuncs[funcName] {
				convData := ConversionData{Signal: signal, convIdx: i}
				needsConvertFunc = append(needsConvertFunc, convData)
			}
			funcName = convertTestName(signal) + strconv.Itoa(i)
			if !existingFuncs[funcName] {
				convData := ConversionData{Signal: signal, convIdx: i}
				needsConvertTestFunc = append(needsConvertTestFunc, convData)
			}
		}
	}
	return needsConvertFunc, needsConvertTestFunc
}

// createConvertTestFunc creates test functions for the conversion functions if they do not exist.
func createConvertTestFunc(tmplData *codegen.TemplateData, outputDir string, needsConvertTestFunc []ConversionData) error {
	convertTestFuncTemplate, err := createConvertTestFuncTemplate(tmplData.PackageName)
	if err != nil {
		return err
	}

	if len(needsConvertTestFunc) == 0 {
		return nil
	}

	convertTestFuncFileName := fmt.Sprintf(convertTestFuncFileNameFormat, strings.ToLower(tmplData.ModelName))
	filePath := filepath.Join(outputDir, convertTestFuncFileName)
	packageName := tmplData.PackageName + "_test"
	err = writeConvertFuncs(needsConvertTestFunc, convertTestFuncTemplate, filePath, packageName)
	if err != nil {
		return err
	}
	return nil
}

func createConvertFuncTemplate() (*template.Template, error) {
	tmpl, err := template.New("convertFuncTemplate").Funcs(template.FuncMap{
		"convertName": convertName,
	}).Parse(convertFuncTemplateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing go struct template: %w", err)
	}
	return tmpl, nil
}

func createConvertTestFuncTemplate(packageNameToTest string) (*template.Template, error) {
	tmpl, err := template.New("convertTestFuncTemplate").Funcs(template.FuncMap{
		"convertName":     func(sig *codegen.SignalInfo) string { return fmt.Sprintf("%s.%s", packageNameToTest, convertName(sig)) },
		"convertTestName": convertTestName,
	}).Parse(convertTestsFuncTemplateStr)
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
		return nil, fmt.Errorf("error reading directory: %w", err)
	}

	for _, d := range list {
		if d.IsDir() || !strings.HasSuffix(d.Name(), ".go") {
			continue
		}
		filename := filepath.Join(outputPath, d.Name())
		err = addFileDeclerations(fset, filename, declaredFunctions)
		if err != nil {
			return nil, err
		}
	}

	return declaredFunctions, nil
}

func addFileDeclerations(fset *token.FileSet, filePath string, declaredFunctions map[string]bool) error {
	src, err := parser.ParseFile(fset, filePath, nil, parser.SkipObjectResolution)
	if err != nil {
		return fmt.Errorf("error parsing file: %w", err)
	}
	for _, decl := range src.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Recv != nil {
			continue
		}
		declaredFunctions[fn.Name.Name] = true
	}
	return nil
}

func convertName(signal *codegen.SignalInfo) string {
	return "To" + signal.GOName
}

func convertTestName(signal *codegen.SignalInfo) string {
	return "Test" + convertName(signal)
}

// ensureFuncFile checks if the convertFunc file exists and creates it if it does not.
// It also writes the package header to the file if it is created.
func ensureFuncFile(convertFuncPath string, packageName string) error {
	_, err := os.Stat(convertFuncPath)
	if err == nil {
		return nil
	}
	if !os.IsNotExist(err) {
		return fmt.Errorf("error checking for %s file: %w", convertFuncPath, err)
	}
	// create the convertFunc file
	funcFile, err := os.Create(filepath.Clean(convertFuncPath))
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

func writeConvertFuncs(needsConvertFunc []ConversionData, tmpl *template.Template, outputPath string, packageName string) error {
	// check if we need to create convertFunc file
	err := ensureFuncFile(outputPath, packageName)
	if err != nil {
		return err
	}

	convertData, err := os.ReadFile(filepath.Clean(outputPath))
	if err != nil {
		return fmt.Errorf("error reading convertFunc file: %w", err)
	}
	convertBuff := bytes.NewBuffer(convertData)
	for _, convData := range needsConvertFunc {
		data := funcTmplData{
			PackageName: packageName,
			Signal:      convData.Signal,
			ConvIdx:     convData.convIdx,
			Conversion:  convData.Signal.Conversions[convData.convIdx],
		}
		if err = tmpl.Execute(convertBuff, &data); err != nil {
			return fmt.Errorf("error executing convertFunc template: %w", err)
		}
	}
	err = codegen.FormatAndWriteToFile(convertBuff.Bytes(), outputPath)
	if err != nil {
		return fmt.Errorf("error formatting and writing to file: %w", err)
	}
	return nil
}
