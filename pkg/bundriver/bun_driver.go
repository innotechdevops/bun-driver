package bundriver

import (
	"github.com/dreamph/dbre"
	"github.com/uptrace/bun"
)

type Config struct {
	Host         string
	Port         int
	DatabaseName string
	User         string
	Pass         string
	PoolOptions  *dbre.DbPoolOptions
	Loc          string
}

type BunDriver interface {
	Connect() *bun.DB
}
