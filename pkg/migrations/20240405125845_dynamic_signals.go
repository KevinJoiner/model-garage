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
	addFunc := func() { goose.AddNamedMigrationContext(filename, upDynamicSignals, downDynamicSignals) }
	registerFuncs = append(registerFuncs, addFunc)
	addFunc()
}

func upDynamicSignals(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, createVehicleSigsTable)
	if err != nil {
		return fmt.Errorf("failed to create dimo table: %w", err)
	}
	return nil
}

func downDynamicSignals(ctx context.Context, tx *sql.Tx) error {
	dropStmt := "DROP TABLE IF EXISTS signal"
	_, err := tx.ExecContext(ctx, dropStmt)
	if err != nil {
		return fmt.Errorf("failed to drop signal table: %w", err)
	}
	return nil
}

const createVehicleSigsTable = `CREATE TABLE IF NOT EXISTS signal(
	TokenID UInt32 COMMENT 'tokenID of this device data.',
    Subject String COMMENT 'subjet of this vehicle data.',
	Timestamp DateTime('UTC') COMMENT 'timestamp of when this data was colllected.',
	Name LowCardinality(String) COMMENT 'name of the signal collected.',
	ValueString Array(String) COMMENT 'value of the signal collected.',
	ValueNumber Array(Float64) COMMENT 'value of the signal collected.',
)
ENGINE = MergeTree()
ORDER BY (TokenID, Timestamp, Name)`
