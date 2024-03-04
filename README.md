# Model Garage

![GitHub license](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
[![GoDoc](https://godoc.org/github.com/KevinJoiner/model-garage?status.svg)](https://godoc.org/github.com/KevinJoiner/model-garage)
[![Go Report Card](https://goreportcard.com/badge/github.com/KevinJoiner/model-garage)](https://goreportcard.com/report/github.com/KevinJoiner/model-garage)

Welcome to the **Model Garage**, a Golang toolkit for managing and working with DIMO models generated from vspec CSV schemas. Model Garage provides the following features:

## Features

1. **Model Creation**: Create models from vspec CSV schemas, represented as Go structs and matching Clickhouse tables.

2. **JSON Conversion**: Easily convert JSON data for your models with automatically generated and customizable conversion functions.

3. **Random Data Generation**: [**Coming Soon**] Quickly generate models populated with random data for testing or sample data.

## Getting Started

1. **Installation**:
   ```bash
   go get github.com/your-username/model-garage
   ```

2. **Import in Your Code**:
   ```go
   import "github.com/your-username/model-garage"
   ```

3. **Usage**:
   Explore the documentation to start using Model Garage in your project.


## Repo structure

### Codegen
The `codegen` directory contains the code generation tool for creating models from vspec CSV schemas. The tool is a standalone application that can be run from the command line.

Example usage:
```bash
go run ./cmd/codegen -output=./pkg/vss -spec=./schema/vss_rel_4.2-DIMO.csv -migrations=./schema/migrations.json -package=vss
package main
```

#### Generation Info
The Model generation is handled by packages in `internal/codegen`. They are responsible for creating Go structs, Clickhouse tables, and conversion functions from the vspec CSV schema and migrations file.

**Vspec Schema** The vspec schema is a CSV file that contains the signal definitions for the model. This schema is generated using vss-tools in this https://github.com/KevinJoiner/DIMO-VSS repository.

**Migrations File** The migrations file is a JSON file that contains the signal definitions that are to be included in the model. This file is manually created. With the following structure:

- **vspecName**: The name of the signal field in the vspec. Only fields specified in the vspec will be included in the model.

- **conversion**: (optional) Details about the conversion from the original data to the vspec field. If not specified, the conversion is assumed to be a direct copy.
	- **originalName**: The original name of the field in the data.

	- **originalType**: (optional) The original data type of the field. If not specified, the original type is assumed to be the same as the vspec type.
##### Generation Process
1. First, the vspec CSV schema and migrations file are parsed.
2. Then a struct is created for each signal in the vpsec schema that is specified in the migrations file. With Clickhouse and JSON tags for each field. The CH and JSON names are the same as the vspec except `.` are replaced with `_`.
3. Next, a Clickhouse table is created for the struct. The table name is the same as the package name. The table is created with the same fields as the struct with corresponding Clickhouse types.
4. Finally, conversion functions are created for each struct. These functions convert the original data in the form of a JSON document to the struct. 

**Conversion Functions**
For each field, a conversion function is created. If a conversion is specified in the migrations file, the conversion function will use the specified conversion. If no conversion is specified, the conversion info function will assume a direct copy. The conversion functions are meant to be overridden with custom logic as needed. When generation is re-run, the conversion functions are not overwritten.