# Model Garage

![GitHub license](https://img.shields.io/badge/license-Apache%202.0-blue.svg)
[![GoDoc](https://godoc.org/github.com/DIMO-Network/model-garage?status.svg)](https://godoc.org/github.com/DIMO-Network/model-garage)
[![Go Report Card](https://goreportcard.com/badge/github.com/DIMO-Network/model-garage)](https://goreportcard.com/report/github.com/DIMO-Network/model-garage)

Welcome to the **Model Garage**, a Golang toolkit for managing and working with DIMO models generated from vspec CSV schemas. Model Garage provides the following features:

## Features

1. **Model Creation**: Create models from vspec CSV schemas, represented as Go structs.
2. **JSON Conversion**: Easily convert JSON data for your models with automatically generated and customizable conversion functions.

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
go run github.com/DIMO-Network/model-garage/cmd/codegen -output=pkg/vss  -generators=all
```

#### Generation Info

The Model generation is handled by packages in `internal/codegen`. They are responsible for creating Go structs, Clickhouse tables, and conversion functions from the vspec CSV schema and definitions file. definitions file is a YAML file that specifies the conversions for each field in the vspec schema. The conversion functions are meant to be overridden with custom logic as needed. When generation is re-run, the conversion functions are not overwritten. Below is an example of a definitions file:

```yaml
# vspecName: The name of the VSpec field in the VSS schema
- vspecName: DIMO.DefinitionID
  # goType: The data type to use for Golang struct.
  # if empty then the type is inferred from the vspec definition
  goType: ""

  # conversion: The mapping of the original data to the VSpec field
  conversion:
    # originalName: The name of the field in the original data
    originalName: data.definitionID

    # originalType: The data type of the field in the original data
    originalType: string

    # isArray: Whether the original field is an array or not
    isArray: false

  # requiredPrivileges: The list of privileges required to access the field
  requiredPrivileges:
    - VEHICLE_NON_LOCATION_DATA
```

##### Generation Process

1. First, the vspec CSV schema and definitions file are parsed.
2. Then a struct is created for each signal in the vpsec schema that is specified in the definitions file. With Clickhouse and JSON tags for each field. The CH and JSON names are the same as the vspec except `.` are replaced with `_`.
3. Next, a Clickhouse table is created for the struct. The table name is the same as the package name. The table is created with the same fields as the struct with corresponding Clickhouse types.
4. Finally, conversion functions are created for each struct. These functions convert the original data in the form of a JSON document to the struct.

**Conversion Functions**
For each field, a conversion function is created. If a conversion is specified in the definitions file, the conversion function will use the specified conversion. If no conversion is specified, the conversion info function will assume a direct copy. The conversion functions are meant to be overridden with custom logic as needed. When generation is re-run, the conversion functions are not overwritten.
