package bundriver

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/dreamph/dbre"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
	"github.com/uptrace/bun/extra/bundebug"
	"log"
)

const MariaDefaultPort = 3306

type MariaConfig Config

type BunMariaDBDriver BunDriver

type mariaDbDriver struct {
	options *MariaConfig
}

func NewMariaDBDriver(options *MariaConfig) BunMariaDBDriver {
	return &mariaDbDriver{
		options: options,
	}
}

func (d *mariaDbDriver) Connect() *bun.DB {
	options := d.options
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=true",
		options.User,
		options.Pass,
		options.Host,
		options.Port,
		options.DatabaseName,
	)
	if options.Loc != "" {
		connection = fmt.Sprintf("%s&loc=%s", connection, options.Loc)
	}

	db, err := sql.Open("mysql", connection)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	ctx := context.Background()
	if err = db.PingContext(ctx); err != nil {
		log.Fatal(err)
		return nil
	}

	dbPool := dbre.DbPoolDefault
	if options.PoolOptions != nil {
		dbPool = options.PoolOptions
	}

	dbre.SetConnectionsPool(db, dbPool)

	bunDB := bun.NewDB(db, mysqldialect.New())
	bunDB.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	return bunDB
}
