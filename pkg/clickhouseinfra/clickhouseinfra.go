// Package clickhouseinfra provides a set of functions to interact with ClickHouse containers.
package clickhouseinfra

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	chmodule "github.com/testcontainers/testcontainers-go/modules/clickhouse"
	"github.com/testcontainers/testcontainers-go/wait"
)

const (
	defaultUser = "default"
	defaultDB   = "dimo"
)

// ColInfo is a struct that holds the column meta information.
type ColInfo struct {
	Name    string
	Type    string
	Comment string
}

// Container is a struct that holds the clickhouse and zookeeper containers.
type Container struct {
	*chmodule.ClickHouseContainer
	ZooKeeperContainer testcontainers.Container
}

// CreateClickHouseContainer function starts and testcontainer for clickhouse.
// The caller is responsible for terminating the container.
func CreateClickHouseContainer(ctx context.Context, userName, password string) (*Container, error) {
	if userName == "" {
		userName = defaultUser
	}
	zkcontainer, zkPort, err := StartZooKeeperContainer(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start zookeeper container: %w", err)
	}
	ipaddr, err := zkcontainer.ContainerIP(ctx)
	if err != nil {
		_ = zkcontainer.Terminate(ctx)
		return nil, fmt.Errorf("failed to get zookeeper container IP: %w", err)
	}
	clickHouseContainer, err := chmodule.RunContainer(ctx,
		testcontainers.WithImage("clickhouse/clickhouse-server:23.3.8.21-alpine"),
		chmodule.WithDatabase(defaultDB),
		chmodule.WithUsername(userName),
		chmodule.WithPassword(password),
		chmodule.WithZookeeper(ipaddr, zkPort),
	)
	if err != nil {
		_ = zkcontainer.Terminate(ctx)
		return nil, fmt.Errorf("failed to start clickhouse container: %w", err)
	}
	return &Container{clickHouseContainer, zkcontainer}, nil
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
func GetClickhouseAsDB(ctx context.Context, container *chmodule.ClickHouseContainer) (*sql.DB, error) {
	host, err := container.ConnectionHost(ctx)
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
	const retries = 3
	for i := 0; i < retries; i++ {
		err = dbConn.Ping()
		if err != nil {
			time.Sleep(500 * time.Millisecond)
			continue
		}
		return dbConn, nil
	}

	return nil, fmt.Errorf("failed to ping clickhouse after %d retries: %w", retries, err)
}

// Terminate function terminates the clickhouse and zookeeper containers.
// If an error occurs, it will be printed to stderr.
func (c *Container) Terminate(ctx context.Context) {
	if err := c.ClickHouseContainer.Terminate(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to terminate clickhouse container: %v", err)
	}
	if err := c.ZooKeeperContainer.Terminate(ctx); err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "failed to terminate clickhouse container: %v", err)
	}
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

// StartZooKeeperContainer function starts a zookeeper container. The caller is responsible for terminating the container.
func StartZooKeeperContainer(ctx context.Context) (testcontainers.Container, string, error) {
	zkPort := nat.Port("2181/tcp")

	zkcontainer, err := testcontainers.GenericContainer(ctx, testcontainers.GenericContainerRequest{
		ContainerRequest: testcontainers.ContainerRequest{
			ExposedPorts: []string{zkPort.Port()},
			Image:        "zookeeper:3.7",
			WaitingFor:   wait.ForListeningPort(zkPort),
		},
		Started: true,
	})
	if err != nil {
		return zkcontainer, "", fmt.Errorf("failed to start zookeeper container: %w", err)
	}
	return zkcontainer, zkPort.Port(), nil
}
