package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	registerFunc := func() { goose.AddNamedMigrationContext(filename, upInitGooseTable, downInitGooseTable) }
	registerFuncs = append(registerFuncs, registerFunc)
	registerFunc()
}

func upInitGooseTable(ctx context.Context, tx *sql.Tx) error {
	const createStmt = `
		CREATE TABLE IF NOT EXISTS goose_db_version ON CLUSTER '{cluster}' as tmp 
		ENGINE = ReplicatedMergeTree('/clickhouse/tables/{cluster}/{database}/{table}/{uuid}', '{replica}')
		ORDER BY date;`

	// This code is executed when the migration is applied.
	upStatements := []string{
		// move goose table to a local tmp table
		"RENAME TABLE goose_db_version TO tmp",
		// make goos table that is replicated across the cluster
		createStmt,
		// copy data from old goose table to new goose table
		"INSERT INTO goose_db_version SELECT * FROM tmp",
		// drop the old goose table
		"DROP TABLE tmp",
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downInitGooseTable(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	downStatements := []string{}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}
