package bundriver

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dreamph/dbre"
	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
)

const PostgresDefaultPort = 5432

// PostgresConfig Additional options specific to PostgreSQL if needed
type PostgresConfig struct {
	Config
	SSLMode     string // disable, require, verify-ca, verify-full
	SearchPath  string // schema search path
	Application string // application_name
}

type PostgresDBDriver DBDriver

type postgresDriver struct {
	options *PostgresConfig
}

func (d *postgresDriver) Connect() *bun.DB {
	options := d.options

	// Build PostgreSQL connection string
	connection := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		options.Host,
		options.Port,
		options.User,
		options.Pass,
		options.DatabaseName,
	)

	// Add optional parameters
	if options.SSLMode != "" {
		connection = fmt.Sprintf("%s sslmode=%s", connection, options.SSLMode)
	}
	if options.SearchPath != "" {
		connection = fmt.Sprintf("%s search_path=%s", connection, options.SearchPath)
	}
	if options.Application != "" {
		connection = fmt.Sprintf("%s application_name=%s", connection, options.Application)
	}
	if options.Loc != "" {
		connection = fmt.Sprintf("%s timezone=%s", connection, options.Loc)
	}

	// Open connection
	db, err := sql.Open("postgres", connection)
	if err != nil {
		log.Fatal(err)
		return nil
	}

	// Test connection
	ctx := context.Background()
	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
		return nil
	}

	// Configure connection pool
	dbPool := dbre.DbPoolDefault
	if options.PoolOptions != nil {
		dbPool = options.PoolOptions
	}
	dbre.SetConnectionsPool(db, dbPool)

	// Initialize Bun with PostgreSQL dialect
	bunDB := bun.NewDB(db, pgdialect.New())

	// Add query debugging hook
	bunDB.AddQueryHook(bundebug.NewQueryHook(
		bundebug.WithVerbose(true),
	))

	return bunDB
}

func NewPostgresDBDriver(options *PostgresConfig) PostgresDBDriver {
	return &postgresDriver{
		options: options,
	}
}
