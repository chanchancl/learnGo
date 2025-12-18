package types

type User struct {
	ID        int
	Username  string
	Email     string
	Age       int
	CreatedAt string // SQLite 的 DATETIME 返回文本
}

type Order struct {
	ID         int     `db:"id"`
	UserID     int     `db:"user_id"`    // 外键关联 users.id
	ProductID  int     `db:"product_id"` // 外键关联 products.id
	Quantity   int     `db:"quantity"`
	TotalPrice float64 `db:"total_price"`
	Status     string  `db:"status"`     // 'pending', 'shipped', 'delivered', 'cancelled'
	OrderDate  string  `db:"order_date"` // 格式: 2023-01-15
	CreatedAt  string  `db:"created_at"` // 完整时间戳
}

type Product struct {
	ID        int     `db:"id"`
	Name      string  `db:"name"`
	Category  string  `db:"category"` // 'electronics', 'clothing', 'books', 'home'
	Price     float64 `db:"price"`
	Stock     int     `db:"stock"`
	Rating    float64 `db:"rating"` // 0.0-5.0
	CreatedAt string  `db:"created_at"`
}
