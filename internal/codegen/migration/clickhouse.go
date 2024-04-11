package migration

import (
	"context"
	"fmt"
	"strings"

	"github.com/DIMO-Network/model-garage/internal/codegen"
	"github.com/DIMO-Network/model-garage/pkg/clickhouseinfra"
	"github.com/DIMO-Network/model-garage/pkg/migrations"
)

// getAlterStatements generates the alter statements for the migration.
// it creates a clickhouse database applies the current migrations and then diffs the current table with the new table.
func getAlterStatements(ctx context.Context, tmplData *codegen.TemplateData) ([]string, []string, error) {
	// start clickhouse test clickhouseinfra.
	chcontainer, err := clickhouseinfra.CreateClickHouseContainer(ctx, "", "")
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create clickhouse container: %w", err)
	}
	defer clickhouseinfra.Terminate(ctx, chcontainer)

	db, err := clickhouseinfra.GetClickhouseAsDB(chcontainer)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get clickhouse db: %w", err)
	}
	err = migrations.RunGoose(ctx, []string{"up", "-v"}, db)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to run db migration: %w", err)
	}

	chConn, err := clickhouseinfra.GetClickHouseAsConn(chcontainer)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get clickhouse connection: %w", err)
	}
	cols, err := clickhouseinfra.GetCurrentCols(ctx, chConn, strings.ToLower(tmplData.ModelName))
	if err != nil {
		return nil, nil, fmt.Errorf("failed to get current cols: %w", err)
	}
	newCols := signalsToColInfo(tmplData.Signals)
	alterStms := calculateStatements(cols, newCols, strings.ToLower(tmplData.ModelName))
	downStms := calculateStatements(newCols, cols, strings.ToLower(tmplData.ModelName))
	return alterStms, downStms, nil
}

func Equal(oldCol, newCol clickhouseinfra.ColInfo) bool {
	return oldCol.Name == newCol.Name && oldCol.Type == newCol.Type && oldCol.Comment == newCol.Comment
}

func signalsToColInfo(signals []*codegen.SignalInfo) []clickhouseinfra.ColInfo {
	colInfos := make([]clickhouseinfra.ColInfo, len(signals))
	for i, sig := range signals {
		colInfos[i] = clickhouseinfra.ColInfo{
			Name:    sig.CHName,
			Type:    sig.CHType(),
			Comment: sig.Desc,
		}
	}
	return colInfos
}

func calculateStatements(oldCols, newCols []clickhouseinfra.ColInfo, dbName string) []string {
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
