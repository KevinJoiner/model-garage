package schema

import (
	_ "embed"
)

//go:embed vss_rel_4.2-DIMO.csv
var VssRel42DIMO []byte

//go:embed definitions.yaml
var Definitions []byte
