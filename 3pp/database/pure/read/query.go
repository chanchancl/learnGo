package read

import (
	"database/sql"
	"fmt"
	"learnGo/3pp/database/pure/types"
	"strings"
)

type User = types.User

// SELECT ? ? , data1, data2

// Read 单条
func GetUser(db *sql.DB, id int) (*User, error) {
	user := &User{}
	err := db.QueryRow(
		"SELECT id, username, email, age, created_at FROM users WHERE id = ?",
		id,
	).Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, fmt.Errorf("用户不存在")
	}
	if err != nil {
		return nil, fmt.Errorf("查询失败: %w", err)
	}
	return user, nil
}

// Read 多条
func ListUsers(db *sql.DB) ([]User, error) {
	rows, err := db.Query("SELECT id, username, email, age, created_at FROM users ORDER BY id")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Age, &u.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return users, nil
}

func GetUsersByIDs(db *sql.DB, ids []int) ([]User, error) {
	if len(ids) == 0 {
		return []User{}, nil
	}

	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}

	query := fmt.Sprintf(`
		SELECT id, username, email, age, created_at
		FROM users
		WHERE id IN (%s)
		ORDER BY id
	`, strings.Join(placeholders, ","))

	rows, err := db.Query(query, args...)
	if err != nil {
		panic(err)
	}

	users := []User{}
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.Age, &user.CreatedAt)
		if err != nil {
			panic(err)
		}
		users = append(users, user)
	}

	return users, nil
}
