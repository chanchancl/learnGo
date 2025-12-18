package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"database/sql"
	"math"

	"github.com/brianvoe/gofakeit/v7"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	_ "modernc.org/sqlite"
)

// BunDB 封装 Bun 数据库操作
type BunDB struct {
	*bun.DB
	ctx context.Context
}

// NewBunDB 创建 Bun 数据库连接
func NewBunDB(dbName string) (*BunDB, error) {
	sqldb, err := sql.Open("sqlite", "./"+dbName)
	if err != nil {
		return nil, fmt.Errorf("无法打开数据库: %w", err)
	}

	// 设置连接池参数
	sqldb.SetMaxOpenConns(25)
	sqldb.SetMaxIdleConns(25)
	sqldb.SetConnMaxLifetime(5 * time.Minute)

	db := bun.NewDB(sqldb, sqlitedialect.New())

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("数据库连接失败: %w", err)
	}

	fmt.Printf("SQLite 数据库打开成功！数据库文件: %s\n", dbName)
	return &BunDB{DB: db, ctx: context.Background()}, nil
}

// CreateTables 创建所有表
func (db *BunDB) CreateTables() error {
	// 创建用户表
	_, err := db.NewCreateTable().
		Model((*User)(nil)).
		IfNotExists().
		Exec(db.ctx)
	if err != nil {
		return fmt.Errorf("创建用户表失败: %w", err)
	}

	// 创建产品表
	_, err = db.NewCreateTable().
		Model((*Product)(nil)).
		IfNotExists().
		Exec(db.ctx)
	if err != nil {
		return fmt.Errorf("创建产品表失败: %w", err)
	}

	// 创建订单表
	_, err = db.NewCreateTable().
		Model((*Order)(nil)).
		IfNotExists().
		Exec(db.ctx)
	if err != nil {
		return fmt.Errorf("创建订单表失败: %w", err)
	}

	fmt.Println("所有表创建成功")
	return nil
}

// ========== CREATE 操作 ==========

// CreateOneUser 创建单个用户
func (db *BunDB) CreateOneUser(username, email string, age int) (int, error) {
	user := &User{
		Username: username,
		Email:    email,
		Age:      age,
	}

	_, err := db.NewInsert().Model(user).Exec(db.ctx)
	if err != nil {
		return 0, fmt.Errorf("插入失败: %w", err)
	}

	fmt.Printf("插入成功，ID: %d\n", user.ID)
	return user.ID, nil
}

// InsertUsers 批量插入用户数据
func (db *BunDB) InsertUsers(count int) error {
	existingCount, err := db.CountAllUsers()
	if err != nil {
		return fmt.Errorf("查询用户数量失败: %w", err)
	}

	if existingCount >= count {
		fmt.Printf("用户数量已足够 (%d ≥ %d)，无需插入\n", existingCount, count)
		return nil
	}

	toInsert := count - existingCount
	successCount := 0

	// 使用事务
	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开始事务失败: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for successCount < toInsert {
		user := &User{
			Username: gofakeit.Username(),
			Email:    gofakeit.Email(),
			Age:      gofakeit.Number(18, 80),
		}

		_, err := tx.NewInsert().Model(user).Exec(db.ctx)
		if err != nil {
			// 唯一性约束冲突时继续
			if strings.Contains(err.Error(), "UNIQUE constraint") {
				continue
			}
			fmt.Printf("插入失败: %v\n", err)
			continue
		}
		successCount++

		if successCount%100 == 0 {
			fmt.Printf("已插入 %d/%d 条用户数据...\n", successCount, toInsert)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	fmt.Printf("成功插入: %d条用户数据\n", successCount)
	return nil
}

// InsertProducts 批量插入产品数据
func (db *BunDB) InsertProducts(count int) error {
	existingCount, err := db.CountAllProducts()
	if err != nil {
		return fmt.Errorf("查询产品数量失败: %w", err)
	}

	if existingCount >= count {
		fmt.Printf("产品数量已足够 (%d ≥ %d)，无需插入\n", existingCount, count)
		return nil
	}

	toInsert := count - existingCount
	successCount := 0
	startTime := time.Now()

	// 预定义的产品类别和类型
	productCategories := []string{
		"electronics", "clothing", "books", "home", "sports",
		"beauty", "toys", "automotive", "grocery", "health"}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开始事务失败: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for successCount < toInsert {
		product := db.generateRandomProduct(productCategories)

		_, err := tx.NewInsert().Model(&product).Exec(db.ctx)
		if err != nil {
			fmt.Printf("插入失败: %v\n", err)
			continue
		}

		successCount++
		if successCount%100 == 0 {
			fmt.Printf("已插入 %d/%d 条产品数据...\n", successCount, toInsert)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	duration := time.Since(startTime)
	rate := float64(successCount) / duration.Seconds()

	fmt.Printf("成功插入: %d条产品数据\n", successCount)
	fmt.Printf("耗时: %.2f秒 (平均 %.2f 条/秒)\n", duration.Seconds(), rate)
	return nil
}

// InsertOrders 批量插入订单数据
func (db *BunDB) InsertOrders(count int) error {
	existingCount, err := db.CountAllOrders()
	if err != nil {
		return fmt.Errorf("获取订单数量失败: %w", err)
	}

	if existingCount >= count {
		return nil
	}

	toInsert := count - existingCount

	// 获取所有用户ID和产品ID
	userIDs, err := db.GetAllUserIDs()
	if err != nil {
		return fmt.Errorf("获取用户ID失败: %w", err)
	}

	productIDs, err := db.GetAllProductIDs()
	if err != nil {
		return fmt.Errorf("获取产品ID失败: %w", err)
	}

	if len(userIDs) == 0 || len(productIDs) == 0 {
		return fmt.Errorf("没有足够的用户或产品数据")
	}

	// 获取产品价格
	productPrices, err := db.GetProductPrices()
	if err != nil {
		return fmt.Errorf("获取产品价格失败: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("开始事务失败: %w", err)
	}

	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	statuses := []string{"pending", "shipped", "delivered", "cancelled"}
	successCount := 0

	for successCount < toInsert {
		userID := userIDs[gofakeit.Number(0, len(userIDs)-1)]
		productID := productIDs[gofakeit.Number(0, len(productIDs)-1)]
		quantity := gofakeit.Number(1, 10)

		price, ok := productPrices[productID]
		if !ok {
			price = gofakeit.Price(10, 500)
		}

		totalPrice := price * float64(quantity)

		// 加权随机选择状态
		statusWeights := []int{15, 30, 50, 5}
		status := statuses[db.weightedRandom(statusWeights)]

		order := &Order{
			UserID:     userID,
			ProductID:  productID,
			Quantity:   quantity,
			TotalPrice: totalPrice,
			Status:     status,
			OrderDate:  time.Now().AddDate(0, 0, -gofakeit.Number(0, 365)),
		}

		_, err := tx.NewInsert().Model(order).Exec(db.ctx)
		if err != nil {
			log.Printf("插入订单失败: %v", err)
			continue
		}

		successCount++
		if successCount%100 == 0 {
			fmt.Printf("已插入 %d/%d 条订单数据...\n", successCount, toInsert)
		}
	}

	if err := tx.Commit(); err != nil {
		return fmt.Errorf("提交事务失败: %w", err)
	}

	fmt.Printf("成功插入: %d条订单数据\n", successCount)
	return nil
}

// ========== READ 操作 ==========

// GetUser 获取单个用户
func (db *BunDB) GetUser(id int) (*User, error) {
	user := new(User)
	err := db.NewSelect().Model(user).Where("id = ?", id).Scan(db.ctx)
	if err != nil {
		return nil, fmt.Errorf("查询用户失败: %w", err)
	}
	return user, nil
}

// GetUsersByIDs 获取多个用户
func (db *BunDB) GetUsersByIDs(ids []int) ([]User, error) {
	var users []User
	err := db.NewSelect().Model(&users).Where("id IN (?)", bun.In(ids)).Scan(db.ctx)
	return users, err
}

// GetAllUsers 获取所有用户
func (db *BunDB) GetAllUsers() ([]User, error) {
	var users []User
	err := db.NewSelect().Model(&users).OrderExpr("id ASC").Scan(db.ctx)
	return users, err
}

// CountAllUsers 统计用户总数
func (db *BunDB) CountAllUsers() (int, error) {
	count, err := db.NewSelect().Model((*User)(nil)).Count(db.ctx)
	return count, err
}

// CountUsersByAge 按年龄范围统计
func (db *BunDB) CountUsersByAge(minAge, maxAge int) (int, error) {
	count, err := db.NewSelect().Model((*User)(nil)).
		Where("age >= ? AND age <= ?", minAge, maxAge).
		Count(db.ctx)
	return count, err
}

// AverageAge 计算平均年龄
func (db *BunDB) AverageAge() (float64, error) {
	var avgAge float64
	err := db.NewSelect().Model((*User)(nil)).
		ColumnExpr("AVG(age) AS avg_age").
		Scan(db.ctx, &avgAge)
	return avgAge, err
}

// GetUserAgeStats 获取年龄统计
func (db *BunDB) GetUserAgeStats() (minAge, maxAge int, err error) {
	var stats struct {
		MinAge int
		MaxAge int
	}

	err = db.NewSelect().Model((*User)(nil)).
		ColumnExpr("MIN(age) AS min_age, MAX(age) AS max_age").
		Scan(db.ctx, &stats)

	return stats.MinAge, stats.MaxAge, err
}

// SumAges 计算年龄总和
func (db *BunDB) SumAges() (int, error) {
	var totalAge int
	err := db.NewSelect().Model((*User)(nil)).
		ColumnExpr("SUM(age) AS total_age").
		Scan(db.ctx, &totalAge)
	return totalAge, err
}

// GroupUsersByAgeRange 按年龄段分组
func (db *BunDB) GroupUsersByAgeRange() ([]AgeGroup, error) {
	var groups []AgeGroup

	err := db.NewSelect().
		ColumnExpr(`
            CASE 
                WHEN age BETWEEN 18 AND 25 THEN '18-25岁'
                WHEN age BETWEEN 26 AND 35 THEN '26-35岁'
                WHEN age BETWEEN 36 AND 45 THEN '36-45岁'
                WHEN age BETWEEN 46 AND 55 THEN '46-55岁'
                WHEN age BETWEEN 56 AND 65 THEN '56-65岁'
                WHEN age > 65 THEN '66岁及以上'
                ELSE '未知'
            END as age_group,
            COUNT(*) as user_count
        `).
		Table("users").
		GroupExpr("1").
		OrderExpr(`
            CASE 
                WHEN age_group = '18-25岁' THEN 1
                WHEN age_group = '26-35岁' THEN 2
                WHEN age_group = '36-45岁' THEN 3
                WHEN age_group = '46-55岁' THEN 4
                WHEN age_group = '56-65岁' THEN 5
                WHEN age_group = '66岁及以上' THEN 6
                ELSE 99
            END
        `).
		Scan(db.ctx, &groups)

	return groups, err
}

// GetUsersPaged 分页获取用户
func (db *BunDB) GetUsersPaged(page, pageSize int) (*PageResult, error) {
	page = max(page, 1)
	pageSize = max(pageSize, 10)
	offset := (page - 1) * pageSize

	var users []User
	err := db.NewSelect().Model(&users).
		OrderExpr("id ASC").
		Limit(pageSize).
		Offset(offset).
		Scan(db.ctx)

	if err != nil {
		return nil, fmt.Errorf("分页查询失败: %w", err)
	}

	total, err := db.CountAllUsers()
	if err != nil {
		return nil, fmt.Errorf("统计总数失败: %w", err)
	}

	totalPages := int(math.Ceil(float64(total) / float64(pageSize)))

	return &PageResult{
		Users:      users,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: totalPages,
	}, nil
}

// SearchUsers 搜索用户
func (db *BunDB) SearchUsers(params SearchParams) (*PageResult, error) {
	query := db.NewSelect().Model((*User)(nil))

	// 构建查询条件
	if params.MinAge > 0 {
		query = query.Where("age >= ?", params.MinAge)
	}
	if params.MaxAge > 0 {
		query = query.Where("age <= ?", params.MaxAge)
	}
	if params.Keyword != "" {
		like := "%" + params.Keyword + "%"
		query = query.Where("(username LIKE ? OR email LIKE ?)", like, like)
	}

	// 分页
	if params.Page > 0 && params.PageSize > 0 {
		offset := (params.Page - 1) * params.PageSize
		query = query.Limit(params.PageSize).Offset(offset)
	}

	var users []User
	err := query.OrderExpr("id ASC").Scan(db.ctx, &users)
	if err != nil {
		return nil, fmt.Errorf("搜索用户失败: %w", err)
	}

	// 查询总数
	countQuery := db.NewSelect().Model((*User)(nil))
	if params.MinAge > 0 {
		countQuery = countQuery.Where("age >= ?", params.MinAge)
	}
	if params.MaxAge > 0 {
		countQuery = countQuery.Where("age <= ?", params.MaxAge)
	}
	if params.Keyword != "" {
		like := "%" + params.Keyword + "%"
		countQuery = countQuery.Where("(username LIKE ? OR email LIKE ?)", like, like)
	}

	total, err := countQuery.Count(db.ctx)
	if err != nil {
		return nil, fmt.Errorf("统计搜索结果失败: %w", err)
	}

	totalPages := 0
	if params.PageSize > 0 {
		totalPages = int(math.Ceil(float64(total) / float64(params.PageSize)))
	}

	return &PageResult{
		Users:      users,
		Total:      total,
		Page:       params.Page,
		PageSize:   params.PageSize,
		TotalPages: totalPages,
	}, nil
}

// ========== UPDATE 操作 ==========

// UpdateUser 更新用户年龄
func (db *BunDB) UpdateUser(id, age int) error {
	user := &User{ID: id, Age: age}
	result, err := db.NewUpdate().Model(user).
		Column("age").
		Where("id = ?", id).
		Exec(db.ctx)

	if err != nil {
		return fmt.Errorf("更新失败: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("未找到 ID 为 %d 的用户", id)
	}

	fmt.Printf("更新成功，影响行数: %d\n", rowsAffected)
	return nil
}

// ========== DELETE 操作 ==========

// DeleteUser 删除用户
func (db *BunDB) DeleteUser(id int) error {
	result, err := db.NewDelete().Model((*User)(nil)).
		Where("id = ?", id).
		Exec(db.ctx)

	if err != nil {
		return fmt.Errorf("删除失败: %w", err)
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		return fmt.Errorf("未找到 ID 为 %d 的用户", id)
	}

	fmt.Printf("删除成功，影响行数: %d\n", rowsAffected)
	return nil
}

// ========== 辅助函数 ==========

// CountAllProducts 统计产品总数
func (db *BunDB) CountAllProducts() (int, error) {
	count, err := db.NewSelect().Model((*Product)(nil)).Count(db.ctx)
	return count, err
}

// CountAllOrders 统计订单总数
func (db *BunDB) CountAllOrders() (int, error) {
	count, err := db.NewSelect().Model((*Order)(nil)).Count(db.ctx)
	return count, err
}

// GetAllUserIDs 获取所有用户ID
func (db *BunDB) GetAllUserIDs() ([]int, error) {
	var ids []int
	err := db.NewSelect().Model((*User)(nil)).Column("id").Scan(db.ctx, &ids)
	return ids, err
}

// GetAllProductIDs 获取所有产品ID
func (db *BunDB) GetAllProductIDs() ([]int, error) {
	var ids []int
	err := db.NewSelect().Model((*Product)(nil)).Column("id").Scan(db.ctx, &ids)
	return ids, err
}

// GetProductPrices 获取产品价格映射
func (db *BunDB) GetProductPrices() (map[int]float64, error) {
	var prices []struct {
		ID    int     `bun:"id"`
		Price float64 `bun:"price"`
	}

	err := db.NewSelect().Model((*Product)(nil)).
		Column("id", "price").
		Scan(db.ctx, &prices)

	if err != nil {
		return nil, err
	}

	priceMap := make(map[int]float64)
	for _, p := range prices {
		priceMap[p.ID] = p.Price
	}
	return priceMap, nil
}

// 生成随机产品
func (db *BunDB) generateRandomProduct(categories []string) Product {
	// 产品类型映射
	productTypes := map[string][]string{
		"electronics": {"Smartphone", "Laptop", "Tablet", "Headphones", "Smart Watch", "Camera", "Speaker", "Monitor"},
		"clothing":    {"T-Shirt", "Jeans", "Jacket", "Dress", "Shoes", "Hat", "Socks", "Sweater"},
		"books":       {"Novel", "Textbook", "Biography", "Cookbook", "Children's Book", "Science Fiction"},
		"home":        {"Lamp", "Chair", "Table", "Bed", "Sofa", "Curtains", "Cookware"},
		"sports":      {"Basketball", "Yoga Mat", "Dumbbells", "Running Shoes", "Tennis Racket"},
		"beauty":      {"Shampoo", "Perfume", "Lipstick", "Face Cream", "Makeup Brush"},
	}

	// 随机选择类别
	category := categories[gofakeit.Number(0, len(categories)-1)]

	// 生成产品名称
	var name string
	if types, ok := productTypes[category]; ok {
		name = types[gofakeit.Number(0, len(types)-1)]

		// 添加品牌
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
		name = gofakeit.ProductName()
	}

	// 为电子产品添加型号
	if category == "electronics" {
		model := fmt.Sprintf("%s-%d", gofakeit.RandomString([]string{"Pro", "Plus", "Ultra", "Lite"}),
			gofakeit.Number(1, 15))
		name = fmt.Sprintf("%s %s", name, model)
	}

	return Product{
		Name:      name,
		Category:  category,
		Price:     gofakeit.Price(1.99, 999.99),
		Stock:     gofakeit.Number(0, 1000),
		Rating:    gofakeit.Float64Range(1.0, 5.0),
		CreatedAt: time.Now().AddDate(0, 0, -gofakeit.Number(0, 365)),
	}
}

// weightedRandom 加权随机选择
func (db *BunDB) weightedRandom(weights []int) int {
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
func (db *BunDB) InsertRelatedData() error {
	fmt.Println("开始插入测试数据...")

	// 1. 插入用户
	fmt.Println("\n1. 插入用户数据...")
	if err := db.InsertUsers(1000); err != nil {
		return fmt.Errorf("插入用户数据失败: %w", err)
	}

	// 2. 插入产品
	fmt.Println("\n2. 插入产品数据...")
	if err := db.InsertProducts(100); err != nil {
		return fmt.Errorf("插入产品数据失败: %w", err)
	}

	// 3. 插入订单
	fmt.Println("\n3. 插入订单数据...")
	if err := db.InsertOrders(100000); err != nil {
		return fmt.Errorf("插入订单数据失败: %w", err)
	}

	fmt.Println("\n4. 数据插入完成！")
	return nil
}
