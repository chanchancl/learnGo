package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type Book struct {
	Title string
	Price int
}

func main() {

	// file: :memory:
	sqldb, err := sql.Open(sqliteshim.ShimName, "file:test.db3")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())

	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	ctx := context.Background()
	_, err = db.NewCreateTable().Model((*Book)(nil)).IfNotExists().Exec(ctx)
	if err != nil {
		panic(err)
	}

	bk := []Book{}
	err = db.NewSelect().Model(&bk).Scan(ctx)
	if err != nil {
		panic(err)
	}

	fmt.Println(bk)
}
