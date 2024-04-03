package migration

import (
	"context"
	"fmt"
	"strings"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/DIMO-Network/model-garage/internal/codegen"
	"github.com/DIMO-Network/model-garage/pkg/container"
	_ "github.com/DIMO-Network/model-garage/pkg/migrations"
	"github.com/pressly/goose/v3"
)

type colInfo struct {
	Name    string
	Type    string
	Comment string
}

// getAlterStatements generates the alter statements for the migration.
// it creates a clickhouse database applies the current migrations and then diffs the current table with the new table.
func getAlterStatements(ctx context.Context, tmplData *codegen.TemplateData) ([]string, []string, error) {
	// start clickhouse test container.
	chcontainer, err := container.CreateClickHouseContainer(ctx, "", "")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create clickhouse container: %w", err)
	}
	defer container.Terminate(ctx, chcontainer)

	db, err := container.GetClickhouseAsDB(chcontainer)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get clickhouse db: %w", err)
	}
	goose.SetDialect("clickhouse")
	err = goose.RunContext(ctx, "up", db, ".")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to apply migrations: %w", err)
	}

	chConn, err := container.GetClickHouseAsConn(chcontainer)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get clickhouse connection: %w", err)
	}
	cols, err := GetCurrentCols(ctx, chConn, strings.ToLower(tmplData.ModelName))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get current cols: %w", err)
	}
	newCols := signalsToColInfo(tmplData.Signals)
	alterStms := calculateStatements(cols, newCols, strings.ToLower(tmplData.ModelName))
	downStms := calculateStatements(newCols, cols, strings.ToLower(tmplData.ModelName))
	return alterStms, downStms, nil
}

func Equal(oldCol, newCol colInfo) bool {
	return oldCol.Name == newCol.Name && oldCol.Type == newCol.Type && oldCol.Comment == newCol.Comment
}

// GetCurrentCols  returns the current columns of the table.
func GetCurrentCols(ctx context.Context, chConn clickhouse.Conn, tableName string) ([]colInfo, error) {
	selectStm := fmt.Sprintf("SELECT name, type, comment  FROM system.columns where table='%s'", tableName)
	rows, err := chConn.Query(ctx, selectStm)
	if err != nil {
		return nil, fmt.Errorf("failed to show table: %w", err)
	}
	defer rows.Close()
	colInfos := []colInfo{}
	count := 0
	for rows.Next() {
		count++
		var info colInfo
		err := rows.Scan(&info.Name, &info.Type, &info.Comment)
		if err != nil {
			return nil, fmt.Errorf("failed to scan table: %w", err)
		}
		colInfos = append(colInfos, info)
	}
	return colInfos, nil
}

func signalsToColInfo(signals []*codegen.SignalInfo) []colInfo {
	colInfos := make([]colInfo, len(signals))
	for i, sig := range signals {
		colInfos[i] = colInfo{
			Name:    sig.CHName,
			Type:    sig.CHType(),
			Comment: sig.Desc,
		}
	}
	return colInfos
}

func calculateStatements(oldCols, newCols []colInfo, dbName string) []string {
	alterStms := []string{}
	i := len(oldCols) - 1
	j := len(newCols) - 1
	for i >= 0 || j >= 0 {
		if i == -1 {
			// add signal to the table.
			altStm := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s COMMENT '%s' FIRST", dbName, newCols[j].Name, newCols[j].Type, newCols[j].Comment)
			alterStms = append(alterStms, altStm)
			j--
			continue
		}
		if j == -1 {
			// drop the current column.
			altStm := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", dbName, oldCols[i].Name)
			alterStms = append(alterStms, altStm)
			i--
			continue
		}

		// if the name  is the same verify they are Equal if yes continue if not alter the column.
		if oldCols[i].Name == newCols[j].Name {
			if !Equal(newCols[j], oldCols[i]) {
				altStm := fmt.Sprintf("ALTER TABLE %s ALTER COLUMN %s %s COMMENT '%s'", dbName, newCols[j].Name, newCols[j].Type, newCols[j].Comment)
				alterStms = append(alterStms, altStm)
			}
			i--
			j--
			continue
		}

		if oldCols[i].Name > newCols[j].Name {
			// drop the current column.
			altStm := fmt.Sprintf("ALTER TABLE %s DROP COLUMN %s", dbName, oldCols[i].Name)
			alterStms = append(alterStms, altStm)
			i--
			continue
		}

		altStm := fmt.Sprintf("ALTER TABLE %s ADD COLUMN %s %s COMMENT '%s' AFTER %s", dbName, newCols[j].Name, newCols[j].Type, newCols[j].Comment, oldCols[i].Name)
		alterStms = append(alterStms, altStm)
		j--
	}

	return alterStms
}
