package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"

	"learnGo/3pp/database/pure/create"
	"learnGo/3pp/database/pure/delete"
	"learnGo/3pp/database/pure/read"
	"learnGo/3pp/database/pure/search"
	"learnGo/3pp/database/pure/types"
	"learnGo/3pp/database/pure/update"
)

const (
	dbName = "learn.db3"
)

type User = types.User

/*
	CREATE TABLE IF NOT EXISTS users (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            username TEXT NOT NULL,
			name type attribute
        )


	// C
	INSERT INTO users(username, email, age) VALUES(?, ?, ?)
	INSERT INFO TABLE(ATTRIBUTE) VALUES(DATA)

	// R
	SELECT id, username, email, age, created_at FROM users WHERE id = ?
	SELECT ATTRIBUTE FROM table [WHERE condition] [ORDER BY] [GROUP]

	// U
	UPDATE users SET age = ? WHERE id = ?
	UPDATE table SET attribute = newvalue WHERE condition

	// D
	DELETE FROM users WHERE id = ?
	DELETE FROM table WHERE condition


	SELECT COUNT(*) FROM users

	use transition to push batch data
*/

const (
	SQL_CREATE_USER = `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			username VARCHAR(50) NOT NULL UNIQUE,
			email VARCHAR(100) NOT NULL UNIQUE,
			age INTEGER CHECK(age >= 0 AND age <= 150),
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`

	SQL_CREATE_PRODUCT = `
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name VARCHAR(100) NOT NULL,
			category VARCHAR(50) NOT NULL,
			price DECIMAL(10, 2) NOT NULL CHECK(price >= 0),
			stock INTEGER NOT NULL DEFAULT 0 CHECK(stock >= 0),
			rating DECIMAL(3, 2) CHECK(rating >= 0 AND rating <= 5),
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`

	SQL_CREATE_ORDER = `
		CREATE TABLE IF NOT EXISTS orders (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			user_id INTEGER NOT NULL,
			product_id INTEGER NOT NULL,
			quantity INTEGER NOT NULL CHECK(quantity > 0),
			total_price DECIMAL(10, 2) NOT NULL CHECK(total_price >= 0),
			status VARCHAR(20) CHECK(status IN ('pending', 'shipped', 'delivered', 'cancelled')),
			order_date DATE NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE,
			FOREIGN KEY (product_id) REFERENCES products(id) ON DELETE CASCADE
		);
	`
)

func connectOrCreateDB(name string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./"+name)
	if err != nil {
		log.Fatal("无法打开数据库看", err)
		return nil, err
	}

	if err = db.Ping(); err != nil {
		log.Fatal("Ping 失败", err)
		return nil, err
	}

	fmt.Printf("SQLite 数据库打开成功！数据库文件: %v\n", name)

	sqls := []string{SQL_CREATE_USER, SQL_CREATE_PRODUCT, SQL_CREATE_ORDER}

	for _, sql := range sqls {
		_, err = db.Exec(sql)
		if err != nil {
			log.Fatal("创建表失败:", err)
			return nil, err
		}
	}

	fmt.Println("用户表创建成功")
	return db, nil
}

func insertData(db *sql.DB) {
	// create 数据
	create.InsertRelatedData(db)
	count, _ := read.CountAllUsers(db)
	if count < 100 {
		create.CreateOneUser(db, "alice", "alice@example.com", 25)
		create.CreateOneUser(db, "bob", "bob@example.com", 30)
	}
	update.UpdateUser(db, 1, 26)
	user, _ := read.GetUser(db, 1)
	fmt.Printf("%#v\n", user)
	delete.DeleteUser(db, 3)
	delete.DeleteUser(db, 4)
}

func basicQuery(db *sql.DB) {
	// 分页查询
	result, _ := search.GetUsersPaged(db, 1, 10)
	fmt.Printf("第 %d 页，共 %d 页，总 %d 条记录\n", result.Page, result.TotalPages, result.Total)
	for _, u := range result.Users {
		fmt.Printf("%2d | %20s | %2d岁\n", u.ID, u.Username, u.Age)
	}

	// 条件搜索
	params := search.SearchParams{
		MinAge:  25,
		MaxAge:  45,
		Keyword: "smith", // 用户名或邮箱包含 smith
	}
	result, err := search.SearchUsers(db, params)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("User搜索结果: 总 %d 条\n", result.Total)
	for _, u := range result.Users {
		fmt.Printf("%2d | %20s | %25s | %2d岁\n", u.ID, u.Username, u.Email, u.Age)
	}

	users, _ := read.GetUsersByIDs(db, []int{1, 2, 5})
	fmt.Printf("共找到 %v 个数据\n", len(users))
	for _, u := range users {
		fmt.Printf("%2d | %20s | %25s | %2d岁\n", u.ID, u.Username, u.Email, u.Age)
	}

	// 根据年龄范围分组
	fmt.Println("\nUser年龄分布")
	ageGroups, err := read.GroupUsersByAgeRange(db)
	if err != nil {
		log.Fatal(err)
	}
	for i := range ageGroups {
		fmt.Printf("%s\t%v\n", ageGroups[i].Group, ageGroups[i].Count)
	}
}

func main() {
	// CRUD
	fmt.Println("开始运行")

	db, err := connectOrCreateDB(dbName)
	if err != nil {
		log.Fatal("无法打开数据库看", err)
	}
	defer db.Close()

	insertData(db)
	basicQuery(db)

	read.Join(db)
	read.Window(db)
}
