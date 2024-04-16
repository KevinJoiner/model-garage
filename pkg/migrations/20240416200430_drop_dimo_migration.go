package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	registerFunc := func() { goose.AddNamedMigrationContext(filename, upCommand20240416200430, downCommand20240416200430) }
	registerFuncs = append(registerFuncs, registerFunc)
	registerFunc()
}

func upCommand20240416200430(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		"DROP TABLE IF EXISTS dimo",
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downCommand20240416200430(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	downStatements := []string{
		createDIMOtableStatement,
	}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}
