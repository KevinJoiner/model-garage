package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	registerFunc := func() { goose.AddNamedMigrationContext(filename, upSignalSourceCol, downSignalSourceCol) }
	registerFuncs = append(registerFuncs, registerFunc)
	registerFunc()
}

func upSignalSourceCol(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		"ALTER TABLE signal_shard ON CLUSTER '{cluster}' ADD COLUMN Source String COMMENT 'source of the signal collected.'",
		"ALTER TABLE signal ON CLUSTER '{cluster}' ADD COLUMN Source String COMMENT 'source of the signal collected.'",
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downSignalSourceCol(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	downStatements := []string{
		"ALTER TABLE signal ON CLUSTER '{cluster}' DROP COLUMN Source",
		"ALTER TABLE signal_shard ON CLUSTER '{cluster}' DROP COLUMN Source",
	}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}
