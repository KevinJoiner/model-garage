// This is custom goose binary with sqlite3 support only.

package main

import (
	"context"
	"log"

	"github.com/DIMO-Network/clickhouse-infra/pkg/migrate"
	"github.com/DIMO-Network/model-garage/pkg/migrations"
)

func main() {
	err := migrate.RunGooseCmd(context.Background(), migrations.RegisterFuncs())
	if err != nil {
		log.Fatal(err)
	}
}
