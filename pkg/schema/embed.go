package schema

import (
	_ "embed"
)

//go:embed spec/vss_rel_4.2-DIMO-*.csv
var vssRel42DIMO string

//go:embed spec/definitions.yaml
var definitionsYAML string

// VssRel42DIMO is the embedded CSV file containing the VSS schema for DIMO.
func VssRel42DIMO() string {
	return vssRel42DIMO
}

// DefinitionsYAML is the embedded YAML file containing the definitions.yaml for the VSS schema.
func DefinitionsYAML() string {
	return definitionsYAML
}
