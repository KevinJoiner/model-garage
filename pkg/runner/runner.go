package runner

import (
	"fmt"
	"io"
	"slices"

	"github.com/DIMO-Network/model-garage/internal/codegen"
	"github.com/DIMO-Network/model-garage/internal/codegen/convert"
	"github.com/DIMO-Network/model-garage/internal/codegen/graphql"
	"github.com/DIMO-Network/model-garage/internal/codegen/migration"
	"github.com/DIMO-Network/model-garage/internal/codegen/model"
)

const (
	// AllGenerator is a constant to run all generators.
	AllGenerator = "all"
	// ModelGenerator is a constant to run the model generator.
	ModelGenerator = "model"
	// ConvertGenerator is a constant to run the convert generator.
	ConvertGenerator = "convert"
	// GraphqlGenerator is a constant to run the graphql generator.
	GraphqlGenerator = "graphql"
	// MigrationGenerator is a constant to run the migration generator.
	MigrationGenerator = "migration"
)

type Config struct {
	PackageName string
	OutputDir   string
	GraphQL     graphql.Config
	Convert     convert.Config
}

// Execute runs the code generation tool.
func Execute(vspecReader, definitionsReader io.Reader, generators []string, cfg Config) error {
	// TODO move params to a config struct.

	if len(generators) == 0 {
		generators = []string{AllGenerator}
	}
	// if none of the generators are selected, return an error.
	switch {
	case slices.Contains(generators, AllGenerator):
	case slices.Contains(generators, ModelGenerator):
	case slices.Contains(generators, ConvertGenerator):
	case slices.Contains(generators, GraphqlGenerator):
	case slices.Contains(generators, MigrationGenerator):
	default:
		return fmt.Errorf("no generator selected")
	}

	err := codegen.EnsureDir(cfg.OutputDir)
	if err != nil {
		return fmt.Errorf("failed to ensure output directory: %w", err)
	}

	tmplData, err := codegen.GetDefinedSignals(vspecReader, definitionsReader)
	if err != nil {
		return fmt.Errorf("failed to get defined signals: %w", err)
	}

	tmplData.PackageName = cfg.PackageName

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, ModelGenerator) {
		err = model.Generate(tmplData, cfg.OutputDir)
		if err != nil {
			return fmt.Errorf("failed to generate model file: %w", err)
		}
	}

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, ConvertGenerator) {
		err = convert.Generate(tmplData, cfg.OutputDir, cfg.Convert)
		if err != nil {
			return fmt.Errorf("failed to generate convert file: %w", err)
		}
	}

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, GraphqlGenerator) {
		err = graphql.Generate(tmplData, cfg.OutputDir, cfg.GraphQL)
		if err != nil {
			return fmt.Errorf("failed to generate graphql file: %w", err)
		}
	}

	if slices.Contains(generators, AllGenerator) || slices.Contains(generators, MigrationGenerator) {
		err = migration.Generate(tmplData, cfg.OutputDir)
		if err != nil {
			return fmt.Errorf("failed to generate migration file: %w", err)
		}
	}
	return nil
}
