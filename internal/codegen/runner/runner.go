package runner

import (
	"fmt"
	"io"
	"slices"

	"github.com/DIMO-Network/model-garage/internal/codegen"
	"github.com/DIMO-Network/model-garage/internal/codegen/clickhouse"
	"github.com/DIMO-Network/model-garage/internal/codegen/convert"
	"github.com/DIMO-Network/model-garage/internal/codegen/graphql"
	"github.com/DIMO-Network/model-garage/internal/codegen/model"
)

const (
	// AllGenerator is a constant to run all generators.
	AllGenerator = "all"
	// ModelGenerator is a constant to run the model generator.
	ModelGenerator = "model"
	// ClickhouseGenerator is a constant to run the clickhouse generator.
	ClickhouseGenerator = "clickhouse"
	// ConvertGenerator is a constant to run the convert generator.
	ConvertGenerator = "convert"
	// GraphqlGenerator is a constant to run the graphql generator.
	GraphqlGenerator = "graphql"
)

// Execute runs the code generation tool.
func Execute(outputDir, packageName string, vspecReader, definitionsReader io.Reader, withTest bool, generators []string) error {
	if len(generators) == 0 {
		generators = []string{AllGenerator}
	}
	err := codegen.EnsureDir(outputDir)
	if err != nil {
		return fmt.Errorf("failed to ensure output directory: %w", err)
	}

	tmplData, err := codegen.GetDefinedSignals(vspecReader, definitionsReader)
	if err != nil {
		return fmt.Errorf("failed to get defined signals: %w", err)
	}

	tmplData.PackageName = packageName

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, ModelGenerator) {
		err = model.Generate(tmplData, outputDir)
		if err != nil {
			return fmt.Errorf("failed to generate model file: %w", err)
		}
	}
	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, ClickhouseGenerator) {
		err = clickhouse.Generate(tmplData, outputDir)
		if err != nil {
			return fmt.Errorf("failed to generate clickhouse file: %w", err)
		}
	}

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, ConvertGenerator) {
		err = convert.Generate(tmplData, outputDir, withTest)
		if err != nil {
			return fmt.Errorf("failed to generate convert file: %w", err)
		}
	}

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, GraphqlGenerator) {
		err = graphql.Generate(tmplData, outputDir)
		if err != nil {
			return fmt.Errorf("failed to generate graphql file: %w", err)
		}
	}
	return nil
}
