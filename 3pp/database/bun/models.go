package main

import (
	"time"

	"github.com/uptrace/bun"
)

type User struct {
	bun.BaseModel `bun:"table:users,alias:u"`

	ID        int       `bun:"id,pk,autoincrement"`
	Username  string    `bun:"username,notnull,unique"`
	Email     string    `bun:"email,notnull,unique"`
	Age       int       `bun:"age,notnull"`
	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
}

type Product struct {
	bun.BaseModel `bun:"table:products,alias:p"`

	ID        int       `bun:"id,pk,autoincrement"`
	Name      string    `bun:"name,notnull"`
	Category  string    `bun:"category,notnull"`
	Price     float64   `bun:"price,notnull"`
	Stock     int       `bun:"stock,notnull,default:0"`
	Rating    float64   `bun:"rating"`
	CreatedAt time.Time `bun:"created_at,default:current_timestamp"`
}

type Order struct {
	bun.BaseModel `bun:"table:orders,alias:o"`

	ID         int       `bun:"id,pk,autoincrement"`
	UserID     int       `bun:"user_id,notnull"`
	ProductID  int       `bun:"product_id,notnull"`
	Quantity   int       `bun:"quantity,notnull"`
	TotalPrice float64   `bun:"total_price,notnull"`
	Status     string    `bun:"status"`
	OrderDate  time.Time `bun:"order_date,notnull"`
	CreatedAt  time.Time `bun:"created_at,default:current_timestamp"`

	// 关联关系
	User    *User    `bun:"rel:belongs-to,join:user_id=id"`
	Product *Product `bun:"rel:belongs-to,join:product_id=id"`
}

// AgeGroup 年龄分组统计
type AgeGroup struct {
	Group string `bun:"age_group"`
	Count int    `bun:"user_count"`
}

// SearchParams 搜索参数
type SearchParams struct {
	MinAge   int
	MaxAge   int
	Keyword  string
	Page     int
	PageSize int
}

// PageResult 分页结果
type PageResult struct {
	Users      []User
	Total      int
	Page       int
	PageSize   int
	TotalPages int
}
