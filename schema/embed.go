package schema

import (
	_ "embed"
)

// VssRel42DIMO is the embedded CSV file containing the VSS schema for DIMO.
//
//go:embed vss_rel_4.2-DIMO.csv
var VssRel42DIMO []byte

// Definitions is the embedded YAML file containing the definitions for the VSS schema.
//
//go:embed definitions.yaml
var Definitions []byte
