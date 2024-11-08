// Package runner is a package that provides a programmatic interface to the code generation tool.
package runner

import (
	"fmt"
	"io"

	"github.com/DIMO-Network/model-garage/internal/generator/convert"
	"github.com/DIMO-Network/model-garage/internal/generator/custom"
	"github.com/DIMO-Network/model-garage/pkg/schema"
)

const (
	// ConvertGenerator is a constant to run the convert generator.
	ConvertGenerator = "convert"
	// CustomDefinitionGenerator is a constant to run the custom generator.
	CustomDefinitionGenerator = "custom-definition"
	// CustomConvertGenerator is a constant to run both the custom and convert generators.
	CustomConvertGenerator = "custom-conversions"
)

// Config is the configuration for the code generation tool.
type Config struct {
	Custom  custom.Config
	Convert convert.Config
}

// Execute runs the code generation tool.
func Execute(vspecReader, definitionsReader io.Reader, generator string, cfg Config) error {
	switch generator {
	case ConvertGenerator:
		tmplData, err := schema.GetDefinedConversionSignals(vspecReader, definitionsReader)
		if err != nil {
			return fmt.Errorf("failed to get defined signals: %w", err)
		}
		err = convert.Generate(tmplData, cfg.Convert)
		if err != nil {
			return fmt.Errorf("failed to generate convert file: %w", err)
		}
	case CustomDefinitionGenerator:
		tmplData, err := schema.GetDefinedSignals(vspecReader, definitionsReader)
		if err != nil {
			return fmt.Errorf("failed to get defined signals: %w", err)
		}
		err = custom.Generate(tmplData, cfg.Custom)
		if err != nil {
			return fmt.Errorf("failed to generate custom file: %w", err)
		}
	case CustomConvertGenerator:
		tmplData, err := schema.GetDefinedConversionSignals(vspecReader, definitionsReader)
		if err != nil {
			return fmt.Errorf("failed to get defined signals: %w", err)
		}
		err = custom.Generate(tmplData, cfg.Custom)
		if err != nil {
			return fmt.Errorf("failed to generate custom conversion file: %w", err)
		}
	default:
		return fmt.Errorf("unknown generator: %s", generator)
	}

	return nil
}
