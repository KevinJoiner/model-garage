package migrations

import "github.com/pressly/goose/v3"

var addFuncs = []func(){}

// SetMigrations sets the migrations for the goose tool.
// this will reset the global migrations to avoid any unwanted migrations registers.
func SetMigrations() {
	goose.ResetGlobalMigrations()
	for _, addFunc := range addFuncs {
		addFunc()
	}
}
