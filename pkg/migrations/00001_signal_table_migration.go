package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	registerFunc := func() { goose.AddNamedMigrationContext(filename, upSignalTable, downSignalTable) }
	registerFuncs = append(registerFuncs, registerFunc)
	registerFunc()
}

func upSignalTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		createSignalStmt,
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downSignalTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	downStatements := []string{
		"DROP TABLE signal",
	}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

const createSignalStmt = `
CREATE TABLE IF NOT EXISTS signal
(
	token_id UInt32 COMMENT 'token_id of this device data.',
	timestamp DateTime64(6, 'UTC') COMMENT 'timestamp of when this data was collected.',
	name LowCardinality(String) COMMENT 'name of the signal collected.',
	source String COMMENT 'source of the signal collected.',
	value_number Float64 COMMENT 'float64 value of the signal collected.',
	value_string String COMMENT 'string value of the signal collected.'
)
ENGINE = ReplacingMergeTree
ORDER BY (token_id, timestamp, name)
`
