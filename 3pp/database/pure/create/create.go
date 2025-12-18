package create

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"learnGo/3pp/database/pure/read"
	"learnGo/3pp/database/pure/types"

	"github.com/brianvoe/gofakeit/v7"
)

type Product = types.Product

// 预定义的产品类别，方便随机选择
var productCategories = []string{
	"electronics", "clothing", "books", "home", "sports",
	"beauty", "toys", "automotive", "grocery", "health",
}

// Create
func CreateOneUser(db *sql.DB, username, email string, age int) (int64, error) {
	result, err := db.Exec(
		"INSERT INTO users(username, email, age) VALUES(?, ?, ?)",
		username, email, age,
	)
	if err != nil {
		return 0, fmt.Errorf("插入失败: %w", err)
	}
	id, _ := result.LastInsertId()
	fmt.Printf("插入成功，ID: %d\n", id)
	return id, nil
}

func InsertUsers(db *sql.DB, count int) error {
	exists, err := read.CountAllUsers(db)
	if err != nil {
		log.Fatal(err)
		return nil
	}
	if exists >= count {
		return nil
	}

	toInsert := count - exists

	tx, err := db.Begin()
	if err != nil {
		log.Printf("开始事务失败: %v", err)
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, _ := tx.Prepare("INSERT INTO users(username, email, age) VALUES(?, ?, ?)")
	successCount := 0
	for successCount < toInsert {
		username := gofakeit.Username() // 随机用户名
		email := gofakeit.Email()       // 随机邮箱
		age := gofakeit.Number(18, 80)  // 18-80岁

		_, err := stmt.Exec(username, email, age)
		if err != nil {
			fmt.Printf("插入失败: %v\n", err)
			continue
		}
		successCount++
	}
	err = tx.Commit()
	if err != nil {
		log.Printf("提交事务失败: %v", err)
		return err
	}
	fmt.Printf("成功插入: %d条数据\n", successCount)
	return nil
}

// InsertProduct 批量插入产品数据
func InsertProduct(db *sql.DB, count int) error {
	// 检查当前已有多少产品
	existingCount, err := countAllProducts(db)
	if err != nil {
		log.Printf("查询产品数量失败: %v", err)
		return err
	}

	if existingCount >= count {
		fmt.Printf("产品数量已足够 (%d ≥ %d)，无需插入\n", existingCount, count)
		return nil
	}

	// 计算需要插入的数量（考虑可能已有部分数据）
	toInsert := count - existingCount

	tx, err := db.Begin()
	if err != nil {
		log.Printf("开始事务失败: %v", err)
		return err
	}

	// 使用 defer 确保事务回滚（如果未提交）
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	// 准备预处理语句
	stmt, err := tx.Prepare(`
		INSERT INTO products 
		(name, category, price, stock, rating, created_at) 
		VALUES(?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		log.Printf("准备预处理语句失败: %v", err)
		return err
	}
	defer stmt.Close()

	successCount := 0
	startTime := time.Now()

	// 批量插入产品数据
	for successCount < toInsert {
		// 使用 gofakeit 生成随机产品数据
		product := generateRandomProduct()

		_, err := stmt.Exec(
			product.Name,
			product.Category,
			product.Price,
			product.Stock,
			product.Rating,
			product.CreatedAt,
		)

		if err != nil {
			log.Printf("插入失败: %v", err)
			continue
		}
		successCount++

		// 每插入100条输出一次进度
		if successCount%100 == 0 {
			fmt.Printf("已插入 %d/%d 条产品数据...\n", successCount, toInsert)
		}
	}

	// 提交事务
	if err := tx.Commit(); err != nil {
		log.Printf("提交事务失败: %v", err)
		return err
	}

	duration := time.Since(startTime)
	rate := float64(successCount) / duration.Seconds()

	fmt.Printf("成功插入: %d条产品数据\n", successCount)
	fmt.Printf("耗时: %.2f秒 (平均 %.2f 条/秒)\n", duration.Seconds(), rate)

	return nil
}

// generateRandomProduct 生成随机产品数据
func generateRandomProduct() Product {
	// 生成产品名称（使用更有意义的组合）
	productTypes := map[string][]string{
		"electronics": {"Smartphone", "Laptop", "Tablet", "Headphones", "Smart Watch", "Camera", "Speaker", "Monitor"},
		"clothing":    {"T-Shirt", "Jeans", "Jacket", "Dress", "Shoes", "Hat", "Socks", "Sweater"},
		"books":       {"Novel", "Textbook", "Biography", "Cookbook", "Children's Book", "Science Fiction"},
		"home":        {"Lamp", "Chair", "Table", "Bed", "Sofa", "Curtains", "Cookware"},
		"sports":      {"Basketball", "Yoga Mat", "Dumbbells", "Running Shoes", "Tennis Racket"},
		"beauty":      {"Shampoo", "Perfume", "Lipstick", "Face Cream", "Makeup Brush"},
	}

	// 随机选择一个类别
	category := productCategories[gofakeit.Number(0, len(productCategories)-1)]

	// 为该类别选择合适的产品类型
	var name string
	if types, ok := productTypes[category]; ok {
		name = types[gofakeit.Number(0, len(types)-1)]
		// 添加品牌或修饰词
		brands := map[string][]string{
			"electronics": {"Apple", "Samsung", "Sony", "Dell", "Bose"},
			"clothing":    {"Nike", "Adidas", "Levi's", "Zara", "H&M"},
			"books":       {"Penguin", "HarperCollins", "Random House", "Simon & Schuster"},
			"home":        {"IKEA", "HomeGoods", "West Elm", "Crate & Barrel"},
			"sports":      {"Wilson", "Nike", "Adidas", "Under Armour"},
		}
		if brandList, hasBrand := brands[category]; hasBrand {
			brand := brandList[gofakeit.Number(0, len(brandList)-1)]
			name = fmt.Sprintf("%s %s", brand, name)
		}
	} else {
		// 默认生成通用产品名
		name = gofakeit.ProductName()
	}

	// 为某些产品添加型号/系列
	if category == "electronics" {
		model := fmt.Sprintf("%s-%d", gofakeit.RandomString([]string{"Pro", "Plus", "Ultra", "Lite"}),
			gofakeit.Number(1, 15))
		name = fmt.Sprintf("%s %s", name, model)
	}

	return Product{
		Name:      name,
		Category:  category,
		Price:     gofakeit.Price(1.99, 999.99),    // 价格在1.99-999.99之间
		Stock:     gofakeit.Number(0, 1000),        // 库存0-1000
		Rating:    gofakeit.Float64Range(1.0, 5.0), // 评分1.0-5.0
		CreatedAt: time.Now().AddDate(0, 0, -gofakeit.Number(0, 365)).Format("2006-01-02 15:04:05"),
	}
}

// InsertOrders 批量插入订单数据（需要关联用户和产品）
func InsertOrders(db *sql.DB, count int) error {
	// 确保有足够的用户和产品

	exists, err := countAllOrders(db)
	if err != nil {
		log.Printf("获取订单数量失败: %v\n", err)
		return nil
	}
	if exists >= count {
		return nil
	}

	toInsert := count - exists

	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	stmt, err := tx.Prepare(`
		INSERT INTO orders 
		(user_id, product_id, quantity, total_price, status, order_date, created_at) 
		VALUES(?, ?, ?, ?, ?, ?, ?)
	`)
	if err != nil {
		return err
	}
	defer stmt.Close()

	// 获取所有用户ID和产品ID
	userIDs, err := getAllUserIDs(db)
	if err != nil {
		return err
	}

	productIDs, err := getAllProductIDs(db)
	if err != nil {
		return err
	}

	// 产品价格缓存
	productPrices, err := getProductPrices(db)
	if err != nil {
		return err
	}

	statuses := []string{"pending", "shipped", "delivered", "cancelled"}
	successCount := 0

	for successCount < toInsert {
		// 随机选择用户和产品
		userID := userIDs[gofakeit.Number(0, len(userIDs)-1)]
		productID := productIDs[gofakeit.Number(0, len(productIDs)-1)]

		// 随机数量
		quantity := gofakeit.Number(1, 10)

		// 计算总价
		price, ok := productPrices[productID]
		if !ok {
			price = gofakeit.Price(10, 500) // 如果缓存中没有，生成一个随机价格
		}
		totalPrice := price * float64(quantity)

		// 随机状态（更多为已发货和已送达）
		statusWeights := []int{15, 30, 50, 5} // pending:15%, shipped:30%, delivered:50%, cancelled:5%
		status := statuses[weightedRandom(statusWeights)]

		// 订单日期（最近一年内）
		orderDate := time.Now().AddDate(0, 0, -gofakeit.Number(0, 365)).Format("2006-01-02")

		_, err := stmt.Exec(
			userID,
			productID,
			quantity,
			totalPrice,
			status,
			orderDate,
			time.Now().Format("2006-01-02 15:04:05"),
		)

		if err != nil {
			log.Printf("插入订单失败: %v", err)
			continue
		}
		successCount++

		if successCount%100 == 0 {
			fmt.Printf("已插入 %d/%d 条订单数据...\n", successCount, count)
		}
	}

	if err := tx.Commit(); err != nil {
		return err
	}

	fmt.Printf("成功插入: %d条订单数据\n", successCount)
	return nil
}

// countAllProducts 统计产品总数
func countAllProducts(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM products").Scan(&count)
	return count, err
}

// 辅助函数
func countAllUsers(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	return count, err
}

func countAllOrders(db *sql.DB) (int, error) {
	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM orders").Scan(&count)
	return count, err
}

func getAllUserIDs(db *sql.DB) ([]int, error) {
	rows, err := db.Query("SELECT id FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func getAllProductIDs(db *sql.DB) ([]int, error) {
	rows, err := db.Query("SELECT id FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var ids []int
	for rows.Next() {
		var id int
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func getProductPrices(db *sql.DB) (map[int]float64, error) {
	rows, err := db.Query("SELECT id, price FROM products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	prices := make(map[int]float64)
	for rows.Next() {
		var id int
		var price float64
		if err := rows.Scan(&id, &price); err != nil {
			return nil, err
		}
		prices[id] = price
	}
	return prices, nil
}

// weightedRandom 加权随机选择
func weightedRandom(weights []int) int {
	total := 0
	for _, w := range weights {
		total += w
	}

	r := gofakeit.Number(0, total-1)
	for i, w := range weights {
		r -= w
		if r < 0 {
			return i
		}
	}
	return len(weights) - 1
}

// InsertRelatedData 插入所有相关数据
func InsertRelatedData(db *sql.DB) error {
	fmt.Println("开始插入测试数据...")

	// 1. 先插入用户
	fmt.Println("\n1. 插入用户数据...")
	if err := InsertUsers(db, 1000); err != nil {
		return err
	}

	// 2. 插入产品
	fmt.Println("\n2. 插入产品数据...")
	if err := InsertProduct(db, 100); err != nil {
		return err
	}

	// 3. 插入订单（需要用户和产品）
	fmt.Println("\n3. 插入订单数据...")
	if err := InsertOrders(db, 100000); err != nil {
		return err
	}

	// 4. 可以继续插入地址、支付、评论等数据...
	fmt.Println("\n4. 数据插入完成！")
	return nil
}
