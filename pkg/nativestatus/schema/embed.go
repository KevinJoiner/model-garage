package schema

import (
	_ "embed"
)

//go:embed native-definitions.yaml
var definitionsYAML string

// DefinitionsYAML is the embedded YAML file containing the definitions.yaml for the VSS schema.
// Deprecated: Use pkg/nativestatus/shema.DefinitionsYAML instead.
func DefinitionsYAML() string {
	return definitionsYAML
}
