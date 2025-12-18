package search

import (
	"database/sql"
	"learnGo/3pp/database/pure/types"
	"math"
)

type User = types.User

type PageResult struct {
	Users      []User
	Total      int
	Page       int
	PageSize   int
	TotalPages int
}

// 搜索第 page 页，每页 pageSize 个元素
func GetUsersPaged(db *sql.DB, page, pageSize int) (*PageResult, error) {
	page = max(page, 1)
	pageSize = max(pageSize, 10)

	offset := (page - 1) * pageSize

	rows, err := db.Query(`
		SELECT id, username, email, age, created_at
		FROM users
		ORDER BY id
		LIMIT ? OFFSET ?
	`, pageSize, offset)
	if err != nil {
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

	var total int
	err = db.QueryRow(`SELECT COUNT(*) FROM users`).Scan(&total)
	if err != nil {
		return nil, err
	}

	return &PageResult{
		Users:      users,
		Total:      total,
		Page:       page,
		PageSize:   pageSize,
		TotalPages: int(math.Ceil(float64(total) / float64(pageSize))),
	}, nil
}
