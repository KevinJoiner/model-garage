package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	registerFunc := func() { goose.AddNamedMigrationContext(filename, upCommand20240411200712, downCommand20240411200712) }
	registerFuncs = append(registerFuncs, registerFunc)
	registerFunc()
}

func upCommand20240411200712(ctx context.Context, tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	upStatements := []string{
		"RENAME TABLE signal TO local_signal ON CLUSTER '{cluster}'",
		clusterSignalCreateStmt,
		distributedSignalCreateStmt,
		"INSERT INTO signal select * FROM clusterAllReplicas('{cluster}', default.local_signal)",
		"DROP TABLE local_signal ON CLUSTER '{cluster}'",
	}
	for _, upStatement := range upStatements {
		_, err := tx.ExecContext(ctx, upStatement)
		if err != nil {
			return err
		}
	}
	return nil
}

func downCommand20240411200712(ctx context.Context, tx *sql.Tx) error {
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

var clusterSignalCreateStmt = `CREATE table signal_shard on CLUSTER '{cluster}' as local_signal
ENGINE = ReplicatedReplacingMergeTree('/clickhouse/tables/{shard}/{database}/{table}', '{replica}') 
ORDER BY (TokenID, Timestamp, Name)`
