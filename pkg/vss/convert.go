package vss

import (
	"fmt"
	"reflect"
	"time"
)

// DIMOToSignals converts a slice of DIMO values to a slice of Signals.
func DIMOToSignals(tokenID uint32, timestamp time.Time, dimos []any) []Signal {
	retSignals := make([]Signal, 0, len(dimos))
	for i, colName := range DimoColNames() {
		if skipCol(colName) {
			continue
		}
		sig := Signal{
			TokenID:   tokenID,
			Timestamp: timestamp,
			Name:      colName,
		}

		switch val := dimos[i].(type) {
		// convert any number type to float64
		case *int, *int8, *int16, *int32, *int64, *uint, *uint8, *uint16, *uint32, *uint64, *float32, *float64:
			num := numToFloat64(val)
			if num == nil {
				continue
			}
			sig.ValueNumber = *num
		case *string:
			if val == nil {
				continue
			}
			sig.ValueString = *val
		case []string:
			if len(val) == 0 {
				continue
			}
			sig.ValueStringArray = val
		case []any:
			if len(val) == 0 {
				continue
			}
			vals := make([]string, len(val))
			for j, v := range val {
				vals[j] = fmt.Sprintf("%v", v)
			}
			sig.ValueStringArray = vals
		default:
			// reflect to see if val is nil
			rVal := reflect.ValueOf(val)
			if rVal.Kind() == reflect.Ptr && rVal.IsNil() {
				continue
			}
			if str, ok := val.(fmt.Stringer); ok {
				sig.ValueString = str.String()
			} else {
				sig.ValueString = fmt.Sprintf("%v", val)
			}
		}
		retSignals = append(retSignals, sig)
	}
	return retSignals
}

// numToFloat64 converts any number type to float64. It does this with a type switch and an individual case for each number type.
func numToFloat64(num any) *float64 {
	switch t := num.(type) {
	case *int:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *int8:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *int16:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *int32:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *int64:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *uint:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *uint8:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *uint16:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *uint32:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *uint64:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *float32:
		if t == nil {
			return nil
		}
		return ref(float64(*t))
	case *float64:
		return t
	default:
		return nil
	}
}

// skipCol returns true if the column should be skipped.
func skipCol(colName string) bool {
	return colName == FieldSubject || colName == FieldTimestamp ||
		colName == FieldType || colName == FieldDefinitionID || colName == FieldSource
}

func ref[T any](t T) *T {
	return &t
}
