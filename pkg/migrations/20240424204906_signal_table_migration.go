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
		createSignalShardStmt,
		distributedSignalCreateStmt,
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
		"DROP TABLE signal ON CLUSTER '{cluster}'",
		"DROP TABLE signal_shard ON CLUSTER '{cluster}'",
	}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

const (
	createSignalShardStmt = `
CREATE TABLE IF NOT EXISTS signal_shard ON CLUSTER '{cluster}'
(
    TokenID UInt32 COMMENT 'tokenID of this device data.',
    Timestamp DateTime64(6, 'UTC') COMMENT 'timestamp of when this data was colllected.',
    Name LowCardinality(String) COMMENT 'name of the signal collected.',
    ValueNumber Float64 COMMENT 'float64 value of the signal collected.',
    ValueString String COMMENT 'string value of the signal collected.'
)
ENGINE = ReplicatedReplacingMergeTree('/clickhouse/tables/{shard}/{database}/{table}/{uuid}', '{replica}') 
ORDER BY (TokenID, Timestamp, Name)
`

	distributedSignalCreateStmt = `
CREATE TABLE IF NOT EXISTS signal ON CLUSTER '{cluster}' AS signal_shard 
ENGINE = Distributed('{cluster}', default, signal_shard, TokenID)
`
)
