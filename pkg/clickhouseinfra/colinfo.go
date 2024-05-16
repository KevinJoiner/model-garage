package clickhouseinfra

import (
	"context"
	"fmt"

	"github.com/ClickHouse/clickhouse-go/v2"
)

// ColInfo is a struct that holds the column meta information.
type ColInfo struct {
	Name    string
	Type    string
	Comment string
}

// GetCurrentCols returns the current columns of the table.
func GetCurrentCols(ctx context.Context, chConn clickhouse.Conn, tableName string) ([]ColInfo, error) {
	selectStm := fmt.Sprintf("SELECT name, type, comment FROM system.columns where table='%s'", tableName)
	rows, err := chConn.Query(ctx, selectStm)
	if err != nil {
		return nil, fmt.Errorf("failed to show table: %w", err)
	}
	defer rows.Close() //nolint // we are not interested in the error here
	colInfos := []ColInfo{}
	count := 0
	for rows.Next() {
		count++
		var info ColInfo
		err := rows.Scan(&info.Name, &info.Type, &info.Comment)
		if err != nil {
			return nil, fmt.Errorf("failed to scan table: %w", err)
		}
		colInfos = append(colInfos, info)
	}
	return colInfos, nil
}
