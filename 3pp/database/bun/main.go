package main

import (
	"fmt"
	"log"
)

func main() {
	fmt.Println("开始运行 Bun ORM 示例")

	// 连接数据库
	db, err := NewBunDB("learn_bun.db3")
	if err != nil {
		log.Fatalf("无法打开数据库: %v", err)
	}
	defer db.Close()

	// 创建表
	if err := db.CreateTables(); err != nil {
		log.Fatalf("创建表失败: %v", err)
	}

	// 插入测试数据
	if err := db.InsertRelatedData(); err != nil {
		log.Printf("插入测试数据失败: %v", err)
	}

	// 基本 CRUD 操作
	runBasicCRUD(db)

	// 查询操作
	runBasicQuery(db)

	// 高级查询
	runAdvancedQuery(db)
}

func runBasicCRUD(db *BunDB) {
	fmt.Println("\n=== 基本 CRUD 操作 ===")

	// 创建用户
	id, err := db.CreateOneUser("alice", "alice@example.com", 25)
	if err != nil {
		log.Printf("创建用户失败: %v", err)
	}

	// 创建另一个用户
	_, err = db.CreateOneUser("bob", "bob@example.com", 30)
	if err != nil {
		log.Printf("创建用户失败: %v", err)
	}

	// 查询用户
	user, err := db.GetUser(id)
	if err != nil {
		log.Printf("查询用户失败: %v", err)
	} else {
		fmt.Printf("查询到用户: %#v\n", user)
	}

	// 更新用户
	if err := db.UpdateUser(id, 26); err != nil {
		log.Printf("更新用户失败: %v", err)
	}

	// 再次查询确认更新
	user, _ = db.GetUser(id)
	fmt.Printf("更新后的用户: %#v\n", user)

	// 删除用户（示例，这里不实际删除重要用户）
	// if err := db.DeleteUser(999); err != nil {
	//     log.Printf("删除用户失败: %v", err)
	// }
}

func runBasicQuery(db *BunDB) {
	fmt.Println("\n=== 基本查询操作 ===")

	// 分页查询
	result, err := db.GetUsersPaged(1, 10)
	if err != nil {
		log.Printf("分页查询失败: %v", err)
	} else {
		fmt.Printf("第 %d 页，共 %d 页，总 %d 条记录\n",
			result.Page, result.TotalPages, result.Total)
		for _, u := range result.Users {
			fmt.Printf("%2d | %20s | %2d岁\n", u.ID, u.Username, u.Age)
		}
	}

	// 条件搜索
	params := SearchParams{
		MinAge:   25,
		MaxAge:   45,
		Keyword:  "smith",
		Page:     1,
		PageSize: 20,
	}

	searchResult, err := db.SearchUsers(params)
	if err != nil {
		log.Printf("搜索用户失败: %v", err)
	} else {
		fmt.Printf("User搜索结果: 总 %d 条\n", searchResult.Total)
		for _, u := range searchResult.Users {
			fmt.Printf("%2d | %20s | %25s | %2d岁\n",
				u.ID, u.Username, u.Email, u.Age)
		}
	}

	// 按ID查询多个用户
	users, err := db.GetUsersByIDs([]int{1, 2, 5})
	if err != nil {
		log.Printf("按ID查询失败: %v", err)
	} else {
		fmt.Printf("共找到 %v 个数据\n", len(users))
		for _, u := range users {
			fmt.Printf("%2d | %20s | %25s | %2d岁\n",
				u.ID, u.Username, u.Email, u.Age)
		}
	}

	// 年龄分组统计
	fmt.Println("\nUser年龄分布")
	ageGroups, err := db.GroupUsersByAgeRange()
	if err != nil {
		log.Printf("年龄分组失败: %v", err)
	} else {
		for i := range ageGroups {
			fmt.Printf("%s\t%v\n", ageGroups[i].Group, ageGroups[i].Count)
		}
	}
}

func runAdvancedQuery(db *BunDB) {
	fmt.Println("\n=== 高级查询操作 ===")

	// 统计总数
	count, err := db.CountAllUsers()
	if err != nil {
		log.Printf("统计用户总数失败: %v", err)
	} else {
		fmt.Printf("用户总数: %d\n", count)
	}

	// 按年龄范围统计
	ageCount, err := db.CountUsersByAge(20, 40)
	if err != nil {
		log.Printf("按年龄统计失败: %v", err)
	} else {
		fmt.Printf("20-40岁用户数: %d\n", ageCount)
	}

	// 平均年龄
	avgAge, err := db.AverageAge()
	if err != nil {
		log.Printf("计算平均年龄失败: %v", err)
	} else {
		fmt.Printf("平均年龄: %.2f\n", avgAge)
	}

	// 年龄统计
	minAge, maxAge, err := db.GetUserAgeStats()
	if err != nil {
		log.Printf("获取年龄统计失败: %v", err)
	} else {
		fmt.Printf("最小年龄: %d, 最大年龄: %d\n", minAge, maxAge)
	}

	// 年龄总和
	totalAge, err := db.SumAges()
	if err != nil {
		log.Printf("计算年龄总和失败: %v", err)
	} else {
		fmt.Printf("年龄总和: %d\n", totalAge)
	}

	// 产品统计
	productCount, err := db.CountAllProducts()
	if err != nil {
		log.Printf("统计产品总数失败: %v", err)
	} else {
		fmt.Printf("产品总数: %d\n", productCount)
	}

	// 订单统计
	orderCount, err := db.CountAllOrders()
	if err != nil {
		log.Printf("统计订单总数失败: %v", err)
	} else {
		fmt.Printf("订单总数: %d\n", orderCount)
	}
}
