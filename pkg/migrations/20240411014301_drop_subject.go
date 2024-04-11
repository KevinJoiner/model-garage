package migrations

import (
	"context"
	"database/sql"
	"runtime"

	"github.com/pressly/goose/v3"
)

func init() {
	_, filename, _, _ := runtime.Caller(0)
	addFunc := func() { goose.AddNamedMigrationContext(filename, upRemoveDropSubject, downRemoveDropSubject) }
	registerFuncs = append(registerFuncs, addFunc)
	addFunc()
}

func upRemoveDropSubject(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "ALTER TABLE signal DROP COLUMN Subject")
	if err != nil {
		return err
	}
	return nil
}

func downRemoveDropSubject(ctx context.Context, tx *sql.Tx) error {
	_, err := tx.ExecContext(ctx, "ALTER TABLE signal ADD COLUMN Subject String After TokenID")
	if err != nil {
		return err
	}
	return nil
}
