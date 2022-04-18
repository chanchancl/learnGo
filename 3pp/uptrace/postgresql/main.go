package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
	"github.com/uptrace/bun/extra/bundebug"
)

func main() {
	dsn := os.Getenv("DSN")
	fmt.Println(dsn)

	sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	if err := sqldb.Ping(); err != nil {
		fmt.Println(err)
		return
	}

	db := bun.NewDB(sqldb, pgdialect.New())
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithEnabled(true), bundebug.WithVerbose(true)))

	QueryModel(db)

	CreateTable(db)
}

func QueryModel(db *bun.DB) {
	items := []struct {
		ID   int    `bun:"id"`
		Name string `bun:"name"`
	}{}
	err := db.NewSelect().Table("items").Model(&items).Scan(context.Background())
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(items)
}

func CreateTable(db *bun.DB) {
	// Model guide
	// https://bun.uptrace.dev/guide/models.html#mapping-tables-to-structs

	users := []struct {
		bun.BaseModel `bun:"table:users,alias:u"`

		ID   int64  `bun:"id,pk,autoincrement"`
		Name string `bun:"name,notnull"`
	}{}
	var errs []error
	defer func() {
		if len(errs) != 0 {
			for i := range errs {
				fmt.Println(errs[i])
			}
		}
	}()

	_, err := db.NewCreateTable().Model(&users).IfNotExists().Exec(context.Background())
	if err != nil {
		errs = append(errs, err)
		return
	}
	cnt, err := db.NewSelect().Model(&users).ScanAndCount(context.Background())
	if err != nil {
		errs = append(errs, err)
		return
	}
	fmt.Println(cnt, users)

	// delete user where id > 1
	_, err = db.NewDelete().Model(&users).Where("u.id > 1").Exec(context.Background())
	if err != nil {
		errs = append(errs, err)
	}

}
