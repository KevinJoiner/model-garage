package clickhouseinfra

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/ClickHouse/ch-go"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/docker/go-connections/nat"
	"github.com/testcontainers/testcontainers-go"
	chmodule "github.com/testcontainers/testcontainers-go/modules/clickhouse"
	"github.com/testcontainers/testcontainers-go/wait"
)

// ColInfo is a struct that holds the column meta information.
type ColInfo struct {
	Name    string
	Type    string
	Comment string
}
type Container struct {
	*chmodule.ClickHouseContainer
	ZooKeeperContainer testcontainers.Container
}

// CreateClickHouseContainer function starts and testcontainer for clickhouse.
// The caller is responsible for terminating the container.
func CreateClickHouseContainer(ctx context.Context, userName, password string) (*Container, error) {
	if userName == "" {
		userName = ch.DefaultUser
	}
	zkcontainer, zkPort, err := StartZooKeeperContainer(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to start zookeeper container: %w", err)
	}
	ipaddr, err := zkcontainer.ContainerIP(ctx)
	if err != nil {
		zkcontainer.Terminate(ctx)
		return nil, fmt.Errorf("failed to get zookeeper container IP: %w", err)
	}
	clickHouseContainer, err := chmodule.RunContainer(ctx,
		testcontainers.WithImage("clickhouse/clickhouse-server:23.3.8.21-alpine"),
		chmodule.WithDatabase(ch.DefaultDatabase),
		chmodule.WithUsername(userName),
		chmodule.WithPassword(password),
		chmodule.WithZookeeper(ipaddr, zkPort),
	)
	if err != nil {
		zkcontainer.Terminate(ctx)
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
	defer rows.Close()
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
