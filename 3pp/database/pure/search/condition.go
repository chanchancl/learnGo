package search

import (
	"database/sql"
	"fmt"
	"strings"
)

type SearchParams struct {
	MinAge  int
	MaxAge  int
	Keyword string
}

// 高级搜索 + 分页
func SearchUsers(db *sql.DB, params SearchParams) (*PageResult, error) {
	var conditions []string
	var args []interface{}

	if params.MaxAge > 0 {
		conditions = append(conditions, "age >= ?")
		args = append(args, params.MinAge)
	}
	if params.MaxAge > 0 {
		conditions = append(conditions, "age <= ?")
		args = append(args, params.MaxAge)
	}
	if params.Keyword != "" {
		conditions = append(conditions, "(username LIKE ? OR email LIKE ?)")
		like := "%s" + params.Keyword + "%"
		args = append(args, like, like)
	}

	whereClause := ""
	if len(conditions) > 0 {
		whereClause = "WHERE " + strings.Join(conditions, " AND ")
	}

	query := fmt.Sprintf(`
		SELECT id, username, email, age, created_at
		FROM users
		%s
		ORDER BY id
	`, whereClause)

	// offset := (params.Page - 1) * params.PageSize
	// args = append(args, params.PageSize, offset)

	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Age, &u.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		panic(err)
		return nil, err
	}

	// 查询总记录数（注意：总记录数要用同样的 WHERE 条件）
	countQuery := fmt.Sprintf("SELECT COUNT(*) FROM users %s", whereClause)
	var total int
	countArgs := args[:] // 去掉 LIMIT 和 OFFSET 参数
	err = db.QueryRow(countQuery, countArgs...).Scan(&total)
	if err != nil {
		panic(err)
		return nil, err
	}

	// totalPages := (total + params.PageSize - 1) / params.PageSize

	return &PageResult{
		Users: users,
		Total: total,
		// Page:       params.Page,
		// PageSize:   params.PageSize,
		// TotalPages: totalPages,
	}, nil
}
