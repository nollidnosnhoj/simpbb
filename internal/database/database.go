package database

import (
	"database/sql"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func NewDb () *bun.DB {
	sqldb, err := sql.Open(sqliteshim.ShimName, "./simpbb.db")	// TODO: Use config to get database name/url
	if err != nil {
		panic(err)
	}
	
	return bun.NewDB(sqldb, sqlitedialect.New())
}