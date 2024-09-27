# Model Garage

![GitHub license](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
[![GoDoc](https://godoc.org/github.com/DIMO-Network/model-garage?status.svg)](https://godoc.org/github.com/DIMO-Network/model-garage)
[![Go Report Card](https://goreportcard.com/badge/github.com/DIMO-Network/model-garage)](https://goreportcard.com/report/github.com/DIMO-Network/model-garage)

Welcome to the **Model Garage**, a Golang toolkit for managing and working with DIMO models generated from vspec CSV schemas. Model Garage provides the following features:

## Features

1. **Model Creation**: Create models from vspec CSV schemas, represented as Go structs.
2. **JSON Conversion**: Easily convert JSON data for your models with automatically generated and customizable conversion functions.

## Signal Definitions

Signal definitions can be found in the [spec package](./pkg/schema/spec/spec.md).
This package is the source of truth for signals used for DIMO Data.

## Migrations

To create a new migration, run the following command:

```bash
Make migration name=<migration_name>
```

This will create a new migration file the given name in the `migrations` directory.
this creation should be used over the goose binary to ensure expected behavior of embedded migrations.

## Repo structure

### Codegen

The `codegen` directory contains the code generation tool for creating models from vspec CSV schemas. The tool is a standalone application that can be run from the command line.

Example usage:

```bash
go run github.com/DIMO-Network/model-garage/cmd/codegen -generators=custom -custom.output-file=./pkg/vss/vehicle-structs.go -custom.template-file=./internal/generator/vehicle.tmpl -custom.format=true
```

```
codegen is a tool to generate code for the model-garage project.
Available generators:
        - custom: Runs a given golang template with pkg/schema.TemplateData data.
        - convert: Generates conversion functions for converting between raw data into signals.Usage:
  -convert.copy-comments
        Copy through comments on conversion functions. Default is false.
  -convert.output-file string
        Output file for the conversion functions. (default "convert-funcs_gen.go")
  -convert.package string
        Name of the package to generate the conversion functions. If empty, the base model name is used.
  -custom.format
        Format the generated file with goimports.
  -custom.output-file string
        Path of the generate gql file (default "custom.txt")
  -custom.template-file string
        Path to the template file. Which is executed with codegen.TemplateData data.
  -definitions string
        Path to the definitions file if empty, the definitions will be used
  -generators string
        Comma separated list of generators to run. Options: convert, custom. (default "all")
  -spec string
        Path to the vspec CSV file if empty, the embedded vspec will be used
```

#### Generation Info

The codegen tool is typically used to create files based on arbitrary signal definitions. The tool reads the signal definitions and custom templates and executes the templates to create the output files.

#### Custom Generator

The custom generator takes in a custom template file and output file. The template file is a Go template that is executed with the signal definitions. The data struct passed into the template is defined by [pkg/schema/signal.go.(TemplateData)](pkg/schema/signal.go)
see [vehicle.tmpl](internal/generator/vehicle.tmpl) for an example template.

#### Convert Generator

The convert generator is a built-in generator that creates conversion functions for each signal. The conversion functions are created based on the signal definitions. The conversion functions are meant to be overridden with custom logic as needed. When generation is re-run, the conversion functions are not overwritten.

## Typical use cases

### Updating mappings

1. Update the signal name to VSS name mappings in [definitions.yaml](./pkg/schema/spec/definitions.yaml).
2. run `make generate`
3. PR and github release

Make the mappings take across our pipeline

1. In the https://github.com/DIMO-Network/benthos-plugin/ repo, update the `go.mod` version for the model-garage dependency.
2. PR and github release
3. In the https://github.com/DIMO-Network/stream-es repo, in `values.yaml` update the container `image.tag` to point to the latest benthos-plugin release commit hash (copy it from release view).
4. Push to main, then go to argo to sync it to prod

### Add signals to DIMO VSS spec

This is when the COVESA standard does not have a signal for something we get from an external integration.

1. Look at this repo: https://github.com/DIMO-Network/VSS/blob/main/overlays/DIMO/dimo.vspec and follow readme there.

```

```
