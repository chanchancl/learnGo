package delete

import (
	"database/sql"
	"fmt"
)

// Delete
func DeleteUser(db *sql.DB, id int) error {
	result, err := db.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return err
	}
	rows, _ := result.RowsAffected()
	if rows == 0 {
		return fmt.Errorf("未找到 ID 为 %d 的用户", id)
	}
	fmt.Printf("删除成功，影响行数: %d\n", rows)
	return nil
}
