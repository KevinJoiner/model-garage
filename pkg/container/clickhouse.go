package container

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/testcontainers/testcontainers-go"
	chmodule "github.com/testcontainers/testcontainers-go/modules/clickhouse"
)

// CreateClickHouseContainer function starts and testcontainer for clickhouse.
// The caller is responsible for terminating the container.
func CreateClickHouseContainer(ctx context.Context, userName, password string) (*chmodule.ClickHouseContainer, error) {
	if userName == "" {
		userName = ch.DefaultUser
	}
	clickHouseContainer, err := chmodule.RunContainer(ctx,
		testcontainers.WithImage("clickhouse/clickhouse-server:23.3.8.21-alpine"),
		chmodule.WithDatabase(ch.DefaultDatabase),
		chmodule.WithUsername(userName),
		chmodule.WithPassword(password),
	)
	if err != nil {
		return clickHouseContainer, fmt.Errorf("failed to run clickhouse container: %w", err)
	}
	return clickHouseContainer, nil
}

// GetClickHouseAsConn function returns a clickhouse.Conn connection which uses native ClickHouse protocol.
func GetClickHouseAsConn(container *chmodule.ClickHouseContainer) (clickhouse.Conn, error) {
	host, err := container.ConnectionHost(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to get clickhouse host: %w", err)
	}
	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{host},
		Auth: clickhouse.Auth{
			Username: container.User,
			Password: container.Password,
			Database: container.DbName,
		},
	})
	if err != nil {
		return nil, fmt.Errorf("failed to open clickhouse connection: %w", err)
	}
	return conn, nil
}

// GetClickhouseAsDB function returns a sql.DB connection which allows interfaceing with the stdlib database/sql package.
func GetClickhouseAsDB(container *chmodule.ClickHouseContainer) (*sql.DB, error) {
	host, err := container.ConnectionHost(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("failed to get clickhouse host: %w", err)
	}
	dbConn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{host},
		Auth: clickhouse.Auth{
			Username: container.User,
			Password: container.Password,
			Database: container.DbName,
		},
	})
	if err := dbConn.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping clickhouse: %w", err)
	}
	return dbConn, nil
}

// Terminate function terminates the clickhouse container.
// If an error occurs, it will be printed to stderr.
func Terminate(ctx context.Context, container *chmodule.ClickHouseContainer) {
	if err := container.Terminate(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to terminate clickhouse container: %v", err)
	}
}
