// Package convert provides a function to generate conversion functions for a vehicle struct.
package convert

import (
	"bytes"
	"cmp"
	_ "embed"
	"fmt"
	"path/filepath"
	"slices"
	"strings"
	"text/template"

	"github.com/DIMO-Network/model-garage/pkg/codegen"
	"github.com/DIMO-Network/model-garage/pkg/schema"
)

const (
	// convertV1FileNameFormat is the name of the Go file that will convert v1 JSON data to the signals.
	convertV1FileNameFormat = "%s-v1-convert.go"

	// convertV2FileNameFormat is the name of the Go file that will convert v2 JSON data to the signals.
	convertV2FileNameFormat = "%s-v2-convert.go"

	// convertFuncFileNameFormat is the name of the Go file that will contain the conversion functions.
	convertFuncFileNameFormat = "%s-convert-funcs.go"
	// convertTestFuncFileNameFormat is the name of the Go file that will contain the conversion test functions.
	convertTestFuncFileNameFormat = "%s-convert-funcs_test.go"
)

type conversionData struct {
	FuncName string
	Signal   *schema.SignalInfo
	convIdx  int
}

//go:embed convertv1.tmpl
var convertV1TemplateStr string

//go:embed convertv2.tmpl
var convertV2TemplateStr string

//go:embed convertFunc.tmpl
var convertFuncTemplateStr string

//go:embed convertTestFunc.tmpl
var convertTestsFuncTemplateStr string

const header = `package %s

// This file is automatically populated with conversion functions for each field of the model struct.
// any conversion functions already defined in this package will not be generated.
// Code generated by model-garage.
`

// Config is the configuration for the conversion generator.
type Config struct {
	// WithTest determines if test functions should be generated.
	WithTest bool
	// CopyComments determines if comments for the conversion functions should be copied through.
	CopyComments bool
}

type funcTmplData struct {
	Signal      *schema.SignalInfo
	FuncName    string
	PackageName string
	Conversion  *schema.ConversionInfo
	DocComment  string
	Body        string
}

type convertTmplData struct {
	*schema.TemplateData
	// Group of conversions by original field name.
	Conversions [][]*singleConversions
}

type singleConversions struct {
	Signal     *schema.SignalInfo
	Conversion *schema.ConversionInfo
}

// Generate creates a conversion functions for each field of a model struct.
// as well as the entire model struct.
func Generate(tmplData *schema.TemplateData, outputDir string, cfg Config) (err error) {
	err = createStructConversion(tmplData, outputDir)
	if err != nil {
		return err
	}

	existingFuncs, err := getDeclaredFunctions(outputDir)
	if err != nil {
		return fmt.Errorf("error getting declared functions: %w", err)
	}

	convertFunc, convertTestFunc := getConversionFunctions(tmplData.Signals)

	err = createConvertFuncs(tmplData, outputDir, cfg.CopyComments, convertFunc, existingFuncs)
	if err != nil {
		return err
	}

	if cfg.WithTest {
		err = createConvertTestFunc(tmplData, outputDir, cfg.CopyComments, convertTestFunc, existingFuncs)
		if err != nil {
			return err
		}
	}

	return nil
}

// createStructConversion creates the conversion function for converting JSON data to a model struct.
func createStructConversion(tmplData *schema.TemplateData, outputDir string) error {
	convV1Tmpl, err := createConvV1Template()
	if err != nil {
		return err
	}

	convV2Tmpl, err := createConvV2Template()
	if err != nil {
		return err
	}

	convSlice := gatherAllConversionsFromSignals(tmplData)
	convTmplData := &convertTmplData{
		TemplateData: tmplData,
		Conversions:  convSlice,
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

// gatherAllConversionsFromSignals gathers all conversions from the signals.
// and returns them in the format [][]*singleConversions.
// Where the outer slice is grouped by the original field name.
// And the inner slice contains a list of conversions for that field and their corresponding signal.
func gatherAllConversionsFromSignals(tmplData *schema.TemplateData) [][]*singleConversions {
	conversions := make(map[string][]*singleConversions, len(tmplData.Signals))
	for _, signal := range tmplData.Signals {
		for _, conv := range signal.Conversions {
			conversions[conv.OriginalName] = append(conversions[conv.OriginalName], &singleConversions{
				Signal:     signal,
				Conversion: conv,
			})
		}
	}
	convSlice := [][]*singleConversions{}
	for _, convs := range conversions {
		convSlice = append(convSlice, convs)
	}
	slices.SortFunc(convSlice, func(i, j []*singleConversions) int {
		return cmp.Compare(i[0].Conversion.OriginalName, j[0].Conversion.OriginalName)
	})
	return convSlice
}
