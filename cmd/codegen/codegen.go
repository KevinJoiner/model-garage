// Package main provides the entrypoint for the code generation tool.
package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/DIMO-Network/model-garage/internal/generator/convert"
	"github.com/DIMO-Network/model-garage/internal/generator/custom"
	"github.com/DIMO-Network/model-garage/pkg/runner"
	"github.com/DIMO-Network/model-garage/pkg/schema"
	"github.com/DIMO-Network/model-garage/pkg/version"
)

func main() {
	// Command-line flags
	printVersion := flag.Bool("version", false, "Print the version of the codegen tool")
	vspecPath := flag.String("spec", "", "Path to the vspec CSV file if empty, the embedded vspec will be used")
	definitionPath := flag.String("definitions", "", "Path to the definitions file if empty, the definitions will be used")
	generators := flag.String("generators", "", "Comma separated list of generators to run. Options: convert, custom.")
	// Convert flags
	copyComments := flag.Bool("convert.copy-comments", false, "Copy through comments on conversion functions. Default is false.")
	convertPackageName := flag.String("convert.package", "", "Name of the package to generate the conversion functions. If empty, the base model name is used.")
	convertOutputFile := flag.String("convert.output-file", convert.DefaultConversionFile, "Output file for the conversion functions.")
	// Custom flags
	customOutFile := flag.String("custom.output-file", custom.DefaultFilePath, "Path of the generate gql file")
	customTemplateFile := flag.String("custom.template-file", "", "Path to the template file. Which is executed with codegen.TemplateData data.")
	customFormat := flag.Bool("custom.format", false, "Format the generated file with goimports.")

	flag.CommandLine.Usage = func() {
		fmt.Fprintf(flag.CommandLine.Output(), `
codegen is a tool to generate code for the model-garage project.
Available generators:
	- custom: Runs a given golang template with pkg/schema.TemplateData data.
	- convert: Generates conversion functions for converting between raw data into signals.`)
		fmt.Fprintf(flag.CommandLine.Output(), "Usage:\n")
		flag.PrintDefaults()
	}

	flag.Parse()
	if *printVersion {
		log.Printf("codegen version: %s", version.GetVersion())
		return
	}

	var vspecReader io.Reader
	if *vspecPath != "" {
		f, err := os.Open(filepath.Clean(*vspecPath))
		if err != nil {
			log.Fatalf("failed to open file: %v", err)
		}
		vspecReader = f
		//nolint:errcheck // we don't care about the error since we are not writing to the file
		defer f.Close()
	} else {
		vspecReader = strings.NewReader(schema.VssRel42DIMO())
	}
	var definitionReader io.Reader
	if *definitionPath != "" {
		f, err := os.Open(filepath.Clean(*definitionPath))
		if err != nil {
			defer log.Fatalf("failed to open file: %v", err)
			return
		}
		definitionReader = f
		//nolint:errcheck // we don't care about the error since we are not writing to the file
		defer f.Close()
	} else {
		definitionReader = strings.NewReader(schema.DefinitionsYAML())
	}
	gens := strings.Split(*generators, ",")

	cfg := runner.Config{
		Custom: custom.Config{
			OutputFile:   *customOutFile,
			TemplateFile: *customTemplateFile,
			Format:       *customFormat,
		},
		Convert: convert.Config{
			CopyComments: *copyComments,
			PackageName:  *convertPackageName,
			OutputFile:   *convertOutputFile,
		},
	}

	err := runner.Execute(vspecReader, definitionReader, gens, cfg)
	if err != nil {
		defer log.Fatal(err)
		return
	}
}
