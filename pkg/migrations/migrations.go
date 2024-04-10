package migrations

import (
	"context"
	"database/sql"
	"embed"
	"fmt"
	"os"
	"time"

	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/pressly/goose/v3"
)

// registerFuncs is a list of functions that register migrations.
// Each migration file should have an init function that appends their register function to this list.
// This is different from the goose registration which is public for all packages.
var registerFuncs = []func(){}

// ClickhouseConfig is the configuration for connecting to a clickhouse database.
type ClickhouseConfig struct {
	Host     string `yaml:"CLICKHOUSE_HOST"`
	Port     string `yaml:"CLICKHOUSE_PORT"`
	User     string `yaml:"CLICKHOUSE_USER"`
	Password string `yaml:"CLICKHOUSE_PASSWORD"`
	DataBase string `yaml:"CLICKHOUSE_DATABASE"`
}

// SetMigrations sets the migrations for the goose tool.
// this will reset the global migrations and FS to avoid any unwanted migrations registers.
func SetMigrations() {
	emptyFs := embed.FS{}
	goose.SetBaseFS(emptyFs)
	goose.ResetGlobalMigrations()
	for _, regFunc := range registerFuncs {
		regFunc()
	}
}

// RunGoose runs the goose command with the provided arguments.
// args should be the command and the arguments to pass to goose.
// eg RunGoose(ctx, []string{"up", "-v"}, chConfig)
func RunGoose(ctx context.Context, gooseArgs []string, chConfig ClickhouseConfig) error {
	if len(gooseArgs) == 0 {
		return fmt.Errorf("command not provided")
	}
	db := GetClickhouseDB(chConfig)
	if err := db.Ping(); err != nil {
		return fmt.Errorf("failed to ping clickhouse: %w", err)
	}
	cmd := gooseArgs[0]
	var args []string
	if len(gooseArgs) > 1 {
		args = os.Args[1:]
	}

	SetMigrations()
	if err := goose.SetDialect("clickhouse"); err != nil {
		return fmt.Errorf("failed to set dialect: %w", err)
	}
	err := goose.RunContext(ctx, cmd, db, ".", args...)
	if err != nil {
		return fmt.Errorf("failed to run goose command: %w", err)
	}
	return nil
}

// GetClickhouseDB returns a sql.DB connection to clickhouse.
func GetClickhouseDB(config ClickhouseConfig) *sql.DB {
	addr := fmt.Sprintf("%s:%s", config.Host, config.Port)
	conn := clickhouse.OpenDB(&clickhouse.Options{
		Addr: []string{addr},
		Auth: clickhouse.Auth{
			Username: config.User,
			Password: config.Password,
			Database: config.DataBase,
		},
		DialTimeout: time.Minute * 30,
	})
	return conn
}
