package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	addFunc := func() { goose.AddNamedMigrationContext(filename, upRemoveValueArray, downRemoveValueArray) }
	registerFuncs = append(registerFuncs, addFunc)
	addFunc()
}

func upRemoveValueArray(ctx context.Context, tx *sql.Tx) error {
	stmts := []string{
		"ALTER TABLE signal RENAME COLUMN ValueNumber TO tmp",
		"ALTER TABLE signal ADD COLUMN ValueNumber Float64 AFTER Name",
		"ALTER TABLE signal COMMENT COLUMN ValueNumber 'float64 value of the signal collected.'",
		"ALTER TABLE signal UPDATE ValueNumber = arrayElement(tmp, 1) WHERE notEmpty(tmp) = 1",
		"ALTER TABLE signal DROP COLUMN tmp",
		"ALTER TABLE signal RENAME COLUMN ValueString TO ValueStringArray",
		"ALTER TABLE signal COMMENT COLUMN ValueStringArray 'string array value of the signal collected.'",
		"ALTER TABLE signal ADD COLUMN ValueString String AFTER ValueNumber",
		"ALTER TABLE signal COMMENT COLUMN ValueString 'string value of the signal collected.'",
		"ALTER TABLE signal UPDATE ValueString = arrayElement(ValueStringArray, 1) WHERE length(ValueStringArray) = 1",
	}
	for _, stmt := range stmts {
		_, err := tx.ExecContext(ctx, stmt)
		if err != nil {
			return err
		}
	}
	return nil
}

func downRemoveValueArray(ctx context.Context, tx *sql.Tx) error {
	stmts := []string{
		"ALTER TABLE signal DROP COLUMN ValueString",
		"ALTER TABLE signal RENAME COLUMN ValueStringArray TO ValueString",
		"ALTER TABLE signal RENAME COLUMN ValueNumber TO tmp",
		"ALTER TABLE signal ADD COLUMN ValueNumber Array(Float64)",
		"ALTER TABLE signal UPDATE ValueNumber = [tmp] WHERE tmp != 0",
		"ALTER TABLE signal DROP COLUMN tmp",
	}
	for _, stmt := range stmts {
		_, err := tx.ExecContext(ctx, stmt)
		if err != nil {
			return err
		}
	}
	return nil
}
