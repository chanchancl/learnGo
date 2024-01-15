package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/go-faker/faker/v4"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type Book struct {
	bun.BaseModel `bun:"table:BooksTable,alias:bk"`

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

	db.NewCreateTable().Model((*Book)(nil)).IfNotExists().Exec(ctx)

	book := Book{
		Title: faker.Username(),
		Price: int(time.Now().Unix()),
	}
	db.NewInsert().Model(&book).Exec(ctx)

	bk := []Book{}
	err = db.NewSelect().Model(&bk).Scan(ctx)
	if err != nil {
		panic(err)
	}

	db.NewDropTable().Model((*Book)(nil)).Exec(ctx)

	fmt.Println(bk)
}
