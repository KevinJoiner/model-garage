package migrations_test

import (
	"context"
	"testing"

	"github.com/DIMO-Network/model-garage/pkg/clickhouseinfra"
	"github.com/DIMO-Network/model-garage/pkg/migrations"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestMigration(t *testing.T) {
	ctx := context.Background()
	chcontainer, err := clickhouseinfra.CreateClickHouseContainer(ctx, "", "")
	require.NoError(t, err, "Failed to create clickhouse container")

	defer chcontainer.Terminate(ctx)

	db, err := clickhouseinfra.GetClickhouseAsDB(chcontainer.ClickHouseContainer)
	require.NoError(t, err, "Failed to get clickhouse db")

	err = migrations.RunGoose(ctx, []string{"up", "-v"}, db)
	require.NoError(t, err, "Failed to run migration")

	conn, err := clickhouseinfra.GetClickHouseAsConn(chcontainer.ClickHouseContainer)
	require.NoError(t, err, "Failed to get clickhouse connection")

	// Iterate over the rows and check the column names
	columns, err := clickhouseinfra.GetCurrentCols(ctx, conn, "signal")
	require.NoError(t, err, "Failed to get current columns")

	expectedColumns := []clickhouseinfra.ColInfo{
		{Name: "TokenID", Type: "UInt32", Comment: "tokenID of this device data."},
		{Name: "Timestamp", Type: "DateTime('UTC')", Comment: "timestamp of when this data was colllected."},
		{Name: "Name", Type: "LowCardinality(String)", Comment: "name of the signal collected."},
		{Name: "ValueNumber", Type: "Float64", Comment: "float64 value of the signal collected."},
		{Name: "ValueString", Type: "String", Comment: "string value of the signal collected."},
		{Name: "ValueStringArray", Type: "Array(String)", Comment: "string array value of the signal collected."},
	}

	// Check if the actual columns match the expected columns
	require.Equal(t, expectedColumns, columns, "Unexpected table columns")

	// Close the DB connection
	err = db.Close()
	assert.NoError(t, err, "Failed to close DB connection")
	err = conn.Close()
	assert.NoError(t, err, "Failed to close clickhouse connection")
}
