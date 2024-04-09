package migrations

import (
	"embed"

	"github.com/pressly/goose/v3"
)

// registerFuncs is a list of functions that register migrations.
// each migration file should have an init function that appends their register function to this list.
// this is different from the goose registration which is public for all packages.
var registerFuncs = []func(){}

// SetMigrations sets the migrations for the goose tool.
// this will reset the global migrations and FS to avoid any unwanted migrations registers.
func SetMigrations() {
	emptyFs := embed.FS{}
	goose.SetBaseFS(emptyFs)
	goose.ResetGlobalMigrations()
	for _, regFunc := range registerFuncs {
		regFunc()
	}
}
