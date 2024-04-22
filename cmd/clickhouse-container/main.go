// Package main is a binary for  creating a test ClickHouse container.
package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/DIMO-Network/model-garage/pkg/clickhouseinfra"
	"github.com/DIMO-Network/model-garage/pkg/migrations"
)

func main() {
	err := run(context.Background())
	if err != nil {
		log.Fatal(err)
	}
}

func run(ctx context.Context) error {
	// Create flags for clickhouse user, password, and port
	user := flag.String("user", "default", "ClickHouse user")
	password := flag.String("password", "default", "ClickHouse password")
	migrate := flag.Bool("migrate", true, "Run migrations")
	flag.Parse()

	chcontainer, err := clickhouseinfra.CreateClickHouseContainer(ctx, *user, *password)
	if err != nil {
		return fmt.Errorf("failed to create clickhouse container: %w", err)
	}
	defer chcontainer.Terminate(ctx)

	if *migrate {
		db, err := clickhouseinfra.GetClickhouseAsDB(ctx, chcontainer.ClickHouseContainer)
		if err != nil {
			return fmt.Errorf("failed to get clickhouse db: %w", err)
		}
		if err := migrations.RunGoose(ctx, []string{"up", "-v"}, db); err != nil {
			return fmt.Errorf("failed to run migration: %w", err)
		}
	}

	host, err := chcontainer.ClickHouseContainer.ConnectionString(ctx)
	if err != nil {
		return fmt.Errorf("failed to get clickhouse host: %w", err)
	}
	fmt.Printf("ClickHouse container is running at: %s\n", host)
	fmt.Println("Waiting for ctrl+c")
	// wait for exit signal to terminate the containers
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
	fmt.Println(" Cya Later!")

	return nil
}
