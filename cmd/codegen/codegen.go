// Package main provides the entrypoint for the code generation tool.
package main

import (
	"bytes"
	"flag"
	"io"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/DIMO-Network/model-garage/internal/codegen/convert"
	"github.com/DIMO-Network/model-garage/internal/codegen/graphql"
	"github.com/DIMO-Network/model-garage/pkg/runner"
	"github.com/DIMO-Network/model-garage/schema"
)

func main() {
	// Command-line flags
	outputDir := flag.String("output", ".", "Output directory for the generated Go file")
	vspecPath := flag.String("spec", "", "Path to the vspec CSV file if empty, the embedded vspec will be used")
	definitionPath := flag.String("definitions", "", "Path to the definitions file if empty, the definitions will be used")
	packageName := flag.String("package", "vspec", "Name of the package to generate")
	generators := flag.String("generators", "all", "Comma separated list of generators to run. Options: all, model, clickhouse, convert, graphql.")
	// Convert flags
	withTest := flag.Bool("convert.with-test", true, "Generate test functions for conversion functions. Default is true.")
	// GQL flags
	gqlOutFile := flag.String("graphql.output-file", "", "Path of the generate gql file that is appened to the outputDir.")
	gqlTemplateFile := flag.String("graphql.template-file", "", "Path to the template file. Which is executed with codegen.TemplateData data.")
	flag.Parse()

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
		vspecReader = bytes.NewReader(schema.VssRel42DIMO)
	}
	var definitionReader io.Reader
	if *definitionPath != "" {
		f, err := os.Open(filepath.Clean(*definitionPath))
		if err != nil {
			log.Fatalf("failed to open file: %v", err)
		}
		definitionReader = f
		//nolint:errcheck // we don't care about the error since we are not writing to the file
		defer f.Close()
	} else {
		definitionReader = bytes.NewReader(schema.Definitions)
	}
	gens := strings.Split(*generators, ",")

	cfg := runner.Config{
		PackageName: *packageName,
		OutputDir:   *outputDir,
		GraphQL: graphql.Config{
			OutputFile:   *gqlOutFile,
			TemplateFile: *gqlTemplateFile,
		},
		Convert: convert.Config{
			WithTest: *withTest,
		},
	}

	err := runner.Execute(vspecReader, definitionReader, gens, cfg)
	if err != nil {
		log.Fatal(err)
	}
}
