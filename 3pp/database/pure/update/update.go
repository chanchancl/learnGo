package update

import (
	"database/sql"
	"fmt"
)

// Update
func UpdateUser(db *sql.DB, id, age int) error {
	// UPDATE users SET age = ? WHERE id = ?
	result, err := db.Exec("UPDATE users SET age = ? WHERE id = ?", age, id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("未找到 ID 为 %d 的用户", id)
	}
	fmt.Printf("更新成功，影响行数: %d\n", rows)
	return nil
}
