package codegen

import (
	"fmt"
	"regexp"
	"strings"
)

const (
	nameCol       = 0
	typeCol       = 1
	dataTypeCol   = 2
	deprecatedCol = 3
	unitCol       = 4
	minCol        = 5
	maxCol        = 6
	descCol       = 7
	colLen        = 8
)

// SignalInfo contains the information for a signal.
// This information comes from the combinations of the spec and migration files.
type SignalInfo struct {
	// From spec CSV
	Name       string
	Type       string
	DataType   string
	Unit       string
	Min        string
	Max        string
	Desc       string
	Deprecated bool

	// Derived
	IsArray    bool
	GOName     string
	BaseGoType string
	BaseCHType string
	CHName     string
	Conversion *ConversionInfo
}

// ConversionInfo contains the conversion information for a field.
type ConversionInfo struct {
	OriginalName string `json:"originalName"`
	OriginalType string `json:"originalType"`
	IsArray      bool   `json:"isArray"`
}

// MigrationInfo contains the migration information for a field.
type MigrationInfo struct {
	IsArray        *bool           `json:"isArray"`
	Conversion     *ConversionInfo `json:"conversion"`
	VspecName      string          `json:"vspecName"`
	ClickHouseType string          `json:"clickHouseType"`
	GoType         string          `json:"goType"`
}

// Migrations is a map of migrations from clickhouse Name to migration info.
type Migrations struct {
	FromName map[string]*MigrationInfo
	Signals  []*SignalInfo
}

// MigratedSignal returns a new slice of signals with the migration information applied.
// excluding signals that are not in the migration file.
func (m *Migrations) MigratedSignal(signal []*SignalInfo) []*SignalInfo {
	sigs := []*SignalInfo{}
	for _, sig := range signal {
		if migration, ok := m.FromName[sig.Name]; ok {
			newSignal := *sig
			newSignal.MergeWithMigration(migration)
			sigs = append(sigs, &newSignal)
		}
	}
	return sigs
}

// NewSignalInfo creates a new SignalInfo from a record from the CSV file.
func NewSignalInfo(record []string) *SignalInfo {
	if len(record) < colLen {
		return nil
	}
	sig := &SignalInfo{
		Name:       record[nameCol],
		Type:       record[typeCol],
		DataType:   record[dataTypeCol],
		Deprecated: record[deprecatedCol] == "true",
		Unit:       record[unitCol],
		Min:        record[minCol],
		Max:        record[maxCol],
		Desc:       record[descCol],
	}
	// arrays are denoted by [] at the end of the type ex uint8[]
	sig.IsArray = strings.HasSuffix(sig.DataType, "[]")
	baseType := sig.DataType
	if sig.IsArray {
		// remove the [] from the type
		baseType = sig.DataType[:len(sig.DataType)-2]
	}
	if baseType != "" {
		//  if this is not a branch type, we can convert it to default golang and clickhouse types
		sig.BaseGoType = goTypeFromVSPEC(baseType)
		sig.BaseCHType = chTypeFromVSPEC(baseType)
	}
	sig.GOName = goName(sig.Name)
	sig.CHName = chName(sig.Name)

	return sig
}

// MergeWithMigration merges the signal with the migration information.
func (s *SignalInfo) MergeWithMigration(migration *MigrationInfo) {
	if migration.ClickHouseType != "" {
		s.BaseCHType = migration.ClickHouseType
	}
	if migration.GoType != "" {
		s.BaseGoType = migration.GoType
	}
	if migration.IsArray != nil {
		s.IsArray = *migration.IsArray
	}
	if migration.Conversion != nil {
		s.Conversion = migration.Conversion
		if s.Conversion.OriginalType == "" {
			s.Conversion.OriginalType = s.GOType()
		}
	}
}

// GOType returns the golang type of the signal.
func (s *SignalInfo) GOType() string {
	if s.IsArray {
		return "[]" + s.BaseGoType
	}
	return s.BaseGoType
}

// CHType returns the clickhouse type of the signal.
func (s *SignalInfo) CHType() string {
	if s.IsArray {
		return "Array(" + s.BaseCHType + ")"
	}
	return s.BaseCHType
}

// CHColType returns the clickhouse proto.Column type of the signal.
func (s *SignalInfo) CHColType() string {
	var builder strings.Builder
	_, _ = builder.WriteString("new(proto.Col")
	switch s.BaseCHType {
	case "UInt8", "Int8", "UInt16", "Int16", "UInt32", "Int32", "UInt64", "Int64", "Float32", "Float64", "Bool", "DateTime":
		_, _ = builder.WriteString(s.BaseCHType)
	case "String":
		_, _ = builder.WriteString("Str")
	default:
		_, _ = builder.WriteString("Auto")
	}
	_, _ = builder.WriteString(")")
	if s.IsArray {
		_, _ = builder.WriteString(".Array()")
	}
	return builder.String()
}

// AppendArg returns the argument to be used in the Append method of the column.
func (s *SignalInfo) AppendArg() string {
	if s.IsArray {
		return fmt.Sprintf("%s{vehicle.%s}", s.GOType(), s.GOName)
	}
	return "vehicle." + s.GOName
}

func goName(name string) string {
	// Remove special characters and ensure PascalCase
	re := regexp.MustCompile("[^a-zA-Z0-9]+")
	return re.ReplaceAllString(name, "")
}

func chName(name string) string {
	return strings.ReplaceAll(name, ".", "_")
}

// goTypeFromVSPEC converts vspec type to golang types.
func goTypeFromVSPEC(dataType string) string {
	switch dataType {
	case "uint8", "int8", "uint16", "int16", "uint32", "int32", "uint64", "int64", "string":
		return dataType
	case "boolean":
		return "bool"
	case "float":
		return "float32"
	case "double":
		return "float64"
	default:
		return "any"
	}
}

// chTypeFromVSPEC converts vspec type to clickhouse types.
func chTypeFromVSPEC(baseType string) string {
	switch baseType {
	case "uint8":
		return "UInt8"
	case "int8":
		return "Int8"
	case "uint16":
		return "UInt16"
	case "int16":
		return "Int16"
	case "uint32":
		return "UInt32"
	case "int32":
		return "Int32"
	case "uint64":
		return "UInt64"
	case "int64":
		return "Int64"
	case "string":
		return "String"
	case "boolean":
		return "Bool"
	case "float":
		return "Float32"
	case "double":
		return "Float64"
	default:
		return "String"
	}
}
