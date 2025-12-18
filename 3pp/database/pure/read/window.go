package read

import (
	"database/sql"
	"fmt"
)

func Window(db *sql.DB) {
	// 通过子查询，查到每种商品在其分类 category 内售卖总数的排行
	// 再通过主查询筛选所有排行第一的商品
	sql := `
	WITH ranked_products AS (
		SELECT
			p.name,
			p.category,
			COUNT(o.id) AS order_count,
			SUM(o.quantity) AS total_sold,
			RANK() OVER (PARTITION BY p.category ORDER BY SUM(o.quantity) DESC) AS rank_in_category
		FROM
			products p
			LEFT JOIN orders o ON p.id = o.product_id
		GROUP BY
			p.id, p.name, p.category
		)
	SELECT
		name,
		category,
		order_count,
		total_sold
	FROM ranked_products
	WHERE rank_in_category = 1
	ORDER BY category;`

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
		var name, category string
		var order_count, total_sold int
		err = rows.Scan(&name, &category, &order_count, &total_sold)
		if err != nil {
			panic(err)
		}
		fmt.Printf("品类 %s 卖的最多的商品是 %s, 总共卖了 %v 单， %v 件\n", category, name, order_count, total_sold)
		count++
	}

}
