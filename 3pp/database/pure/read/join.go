package read

import (
	"database/sql"
	"fmt"
	"learnGo/3pp/database/pure/types"
)

func Join(db *sql.DB) {
	sql := `
	SELECT 
		u.username,
		p.name as product_name,
		o.total_price,
		o.quantity
	FROM users u
	INNER JOIN orders o ON u.id = o.user_id
	INNER JOIN products p ON o.product_id = p.id
	WHERE o.status = 'delivered';`

	rows, err := db.Query(sql)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	count := 0

	for rows.Next() {
		if count > 10 {
			break
		}
		var user types.User
		var order types.Order
		var product types.Product
		err = rows.Scan(&user.Username, &product.Name, &order.TotalPrice, &order.Quantity)
		if err != nil {
			panic(err)
		}
		fmt.Printf("%v 买了%v件%v商品,总共花了%v钱\n", user.Username, order.Quantity, product.Name, order.TotalPrice)
		count++
	}
}
