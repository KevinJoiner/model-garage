//go:generate go run ./internal/codegen/cmd -output=./pkg/vss -spec=./schema/vss_rel_4.2-DIMO.csv -migrations=./schema/migrations.json -package=vss
package main
