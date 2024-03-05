// Package main provides the entrypoint for the code generation tool.
package main

import (
	"flag"
	"log"

	"github.com/KevinJoiner/model-garage/internal/codegen"
	"github.com/KevinJoiner/model-garage/internal/codegen/clickhouse"
	"github.com/KevinJoiner/model-garage/internal/codegen/convert"
	"github.com/KevinJoiner/model-garage/internal/codegen/model"
)

func main() {
	// Command-line flags
	outputDir := flag.String("output", ".", "Output directory for the generated Go file")
	vspecPath := flag.String("spec", "./vspec.csv", "Path to the vspec CSV file")
	migrationPath := flag.String("migrations", "./migrations.json", "Path to the migrations JSON file")
	packageName := flag.String("package", "vspec", "Name of the package to generate")
	withTest := flag.Bool("convert.with-test", true, "Generate test functions for conversion functions. Default is true.")
	flag.Parse()

	err := codegen.EnsureDir(*outputDir)
	if err != nil {
		log.Fatal(err)
	}

	signals, err := codegen.GetMigratedSignals(*vspecPath, *migrationPath)
	if err != nil {
		log.Fatal(err)
	}

	tmplData := codegen.TemplateData{
		PackageName: *packageName,
		Signals:     signals,
	}

	err = model.Generate(&tmplData, *outputDir)
	if err != nil {
		log.Fatalf("failed to generate model: %v", err)
	}

	err = clickhouse.Generate(&tmplData, *outputDir)
	if err != nil {
		log.Fatalf("failed to generate ClickHouse file: %v", err)
	}

	err = convert.Generate(&tmplData, *outputDir, *withTest)
	if err != nil {
		log.Fatalf("failed to generate convert file: %v", err)
	}
}
