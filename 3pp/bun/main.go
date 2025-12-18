package main

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/go-faker/faker/v4"

	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
	"github.com/uptrace/bun/extra/bundebug"
)

type Book struct {
	bun.BaseModel `bun:"table:BooksTable,alias:bk"`

	ID    int `bun:",pk"`
	Title string
	Price int
	Tag   *Tag `bun:"rel:has-one,join:id=book_id"`
}

type Tag struct {
	bun.BaseModel `bun:"table:Tag,alias:tag"`

	BookId  int `bun:",pk"`
	Type    int
	Length  int
	TagList []string
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
		Tag: &Tag{
			Type:    10,
			Length:  3,
			TagList: []string{"a", "b", "c"},
		},
	}
	db.NewInsert().Model(&book).Exec(ctx)

	bk := []Book{}
	err = db.NewSelect().Model(&bk).Scan(ctx)
	if err != nil {
		panic(err)
	}

	db.NewDropTable().Model((*Book)(nil)).Exec(ctx)

	fmt.Printf("%#v\n", bk)
	os.Remove("./test.db3")
}
