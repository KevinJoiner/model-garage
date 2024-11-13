// Package schema information for converting native dimo status objects
package schema

import (
	_ "embed"
)

//go:embed native-definitions.yaml
var definitionsYAML string

// DefinitionsYAML is the embedded YAML file containing the conversion-definitions.yaml for a VSS schema.
func DefinitionsYAML() string {
	return definitionsYAML
}
