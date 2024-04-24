package migrations

import (
	"context"
	"database/sql"
	"fmt"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	registerFunc := func() { goose.AddNamedMigrationContext(filename, upCommand20240423004039, downCommand20240423004039) }
	registerFuncs = append(registerFuncs, registerFunc)
	registerFunc()
}

func upCommand20240423004039(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		"DROP TABLE default.signal ON CLUSTER '{cluster}'",
		"RENAME TABLE default.signal_shard TO default.signal_shard_old ON CLUSTER '{cluster}'",
		createSignalShardStmt,
		"INSERT INTO default.signal_shard SELECT TokenID, Timestamp, Name, ValueNumber, ValueString FROM default.signal_shard_old",
		distributedSignalCreateStmt,
		"DROP TABLE default.signal_shard_old ON CLUSTER '{cluster}'",
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return fmt.Errorf("failed to execute statment %s: %w", upStatement, err)
		}
	}
	return nil
}

func downCommand20240423004039(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	downStatements := []string{
		"ALTER TABLE signal MODIFY COLUMN Timestamp Datetime('UTC')",
	}
	for _, downStatement := range downStatements {
		_, err := tx.ExecContext(ctx, downStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

var (
	createSignalShardStmt = `
CREATE TABLE default.signal_shard ON CLUSTER '{cluster}'
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
CREATE TABLE signal ON CLUSTER '{cluster}' AS signal_shard 
ENGINE = Distributed('{cluster}', default, signal_shard, TokenID)
`
)
