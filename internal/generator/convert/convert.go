// Package convert provides a function to generate conversion functions for a vehicle struct.
package convert

import (
	"bytes"
	_ "embed"
	"fmt"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/pkg/codegen"
	"github.com/DIMO-Network/model-garage/pkg/schema"
)

const (
	// convertV1FileNameFormat is the name of the Go file that will convert v1 JSON data to the signals.
	convertV1FileNameFormat = "%s-v1-convert_gen.go"

	// convertV2FileNameFormat is the name of the Go file that will convert v2 JSON data to the signals.
	convertV2FileNameFormat = "%s-v2-convert_gen.go"

	// convertFuncFileNameFormat is the name of the Go file that will contain the conversion functions.
	convertFuncFileNameFormat = "%s-convert-funcs_gen.go"
)

//go:embed convertv1.tmpl
var convertV1TemplateStr string

//go:embed convertv2.tmpl
var convertV2TemplateStr string

//go:embed convertFunc.tmpl
var convertFuncTemplateStr string

const header = `// Code generated by github.com/DIMO-Network/model-garage.
package %s

// This file is automatically populated with conversion functions for each field of the model struct.
// any conversion functions already defined in this package will be coppied through.
// note: DO NOT mutate the orginalDoc parameter which is shared between all conversion functions.
`

// Config is the configuration for the conversion generator.
type Config struct {
	// CopyComments determines if comments for the conversion functions should be copied through.
	CopyComments bool
	// PackageName is the name of the package to generate the conversion functions.
	// This is separate from the model package name.
	// if empty, the model package name is used.
	PackageName string
	// OutputDir is the output directory for the generated conversion files.
	// if empty, the base output directory is used.
	OutputDir string
}

// funcTmplData contains the data to be used during template execution for writing a single conversion function.
type funcTmplData struct {
	// Signal is the signal that we are converting to.
	Signal *schema.SignalInfo
	// Conversion is the information about the signal we are converting from.
	Conversion *schema.ConversionInfo
	// FuncName is the name of the conversion function.
	FuncName string
	// DocComment is the original doc comment for the conversion function if it exists.
	DocComment string
	// Body of the original conversion function if it exists.
	Body string
}

// convertTmplData contains the data to be used during template execution for writing the conversion functions.
type convertTmplData struct {
	*schema.TemplateData
	ModelPackagePrefix string
}

// Generate creates a conversion functions for each field of a model struct.
// as well as the entire model struct.
func Generate(tmplData *schema.TemplateData, outputDir string, cfg Config) (err error) {
	modelPackagePrefix := ""
	if cfg.PackageName != "" {
		modelPackagePrefix = tmplData.PackageName + "."
		tmplData.PackageName = cfg.PackageName
	}
	if cfg.OutputDir != "" {
		outputDir = cfg.OutputDir
	}

	convertFunc := getConversionFunctions(tmplData.Signals)

	err = createStructConversion(tmplData, convertFunc, outputDir, modelPackagePrefix)
	if err != nil {
		return err
	}

	existingFuncs, err := GetDeclaredFunctions(outputDir)
	if err != nil {
		return fmt.Errorf("error getting declared functions: %w", err)
	}

	err = createConvertFuncs(tmplData, outputDir, cfg.CopyComments, convertFunc, existingFuncs)
	if err != nil {
		return err
	}

	return nil
}

// createStructConversion creates the conversion function for converting JSON data to a model struct.
func createStructConversion(tmplData *schema.TemplateData, conversionFunc []funcTmplData, outputDir, modelPackagePrefix string) error {
	convV1Tmpl, err := createConvV1Template()
	if err != nil {
		return err
	}

	convV2Tmpl, err := createConvV2Template()
	if err != nil {
		return err
	}

	// convSlice := gatherAllConversionsFromSignals3(conversionFunc)
	convTmplData := &convertTmplData{
		TemplateData: tmplData,
		// Conversions:        convSlice,
		ModelPackagePrefix: modelPackagePrefix,
	}

	var outBuf bytes.Buffer
	if err = convV1Tmpl.Execute(&outBuf, &convTmplData); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}
	convertV1FileName := fmt.Sprintf(convertV1FileNameFormat, strings.ToLower(tmplData.ModelName))
	goOutputPath := filepath.Join(outputDir, convertV1FileName)
	// format and write the go file.
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), goOutputPath)
	if err != nil {
		return fmt.Errorf("error formatting and writing to file: %w", err)
	}

	outBuf.Reset()
	if err = convV2Tmpl.Execute(&outBuf, &convTmplData); err != nil {
		return fmt.Errorf("error executing template: %w", err)
	}
	convertV2FileName := fmt.Sprintf(convertV2FileNameFormat, strings.ToLower(tmplData.ModelName))
	goOutputPath = filepath.Join(outputDir, convertV2FileName)
	err = codegen.FormatAndWriteToFile(outBuf.Bytes(), goOutputPath)
	if err != nil {
		return fmt.Errorf("error formatting and writing to file: %w", err)
	}

	return nil
}

func createConvV1Template() (*template.Template, error) {
	tmpl, err := template.New("convertV1Template").Funcs(template.FuncMap{
		"convertName": convertName,
		"lower":       strings.ToLower,
	}).Parse(convertV1TemplateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing conversion v1 template: %w", err)
	}
	return tmpl, nil
}

func createConvV2Template() (*template.Template, error) {
	tmpl, err := template.New("convertV2Template").Funcs(template.FuncMap{
		"convertName": convertName,
		"lower":       strings.ToLower,
	}).Parse(convertV2TemplateStr)
	if err != nil {
		return nil, fmt.Errorf("error parsing conversion v2 template: %w", err)
	}
	return tmpl, nil
}
