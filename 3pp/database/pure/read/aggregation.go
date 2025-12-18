package read

import (
	"database/sql"
)

func CountAllUsers(db *sql.DB) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM users"
	err := db.QueryRow(query).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}

// 统计满足条件的用户数
func CountUsersByAge(db *sql.DB, minAge, maxAge int) (int, error) {
	var count int
	query := "SELECT COUNT(*) FROM users WHERE age >= ? AND age <="
	err := db.QueryRow(query, minAge, maxAge).Scan(&count)
	return count, err
}

func AverageAge(db *sql.DB) (float64, error) {
	var avgAge float64
	query := "SELECT AVG(age) FROM users"
	err := db.QueryRow(query).Scan(&avgAge)
	return avgAge, err
}

func AverageAgeNullable(db *sql.DB) (*float64, error) {
	var avgAge sql.NullFloat64
	query := "SELECT AVG(age) FROM users"
	err := db.QueryRow(query).Scan(&avgAge)
	if err != nil {
		return nil, err
	}
	if avgAge.Valid {
		return &avgAge.Float64, nil
	}
	return nil, nil // 没有数据
}

func GetUserAgeStats(db *sql.DB) (minAge, maxAge int, err error) {
	query := "SELECT MIN(age), MAX(age) FROM users"
	err = db.QueryRow(query).Scan(&minAge, &maxAge)
	return
}

func SumAges(db *sql.DB) (int, error) {
	var totalAge int
	query := "SELECT SUM(age) FROM users"
	err := db.QueryRow(query).Scan(&totalAge)
	return totalAge, err
}

type AgeGroup struct {
	Group string // 如 "18-25岁"
	Count int
}

func GroupUsersByAgeRange(db *sql.DB) ([]AgeGroup, error) {
	query := `
    SELECT 
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
    FROM users
    GROUP BY age_group
    ORDER BY
		CASE age_group
			WHEN '18-25岁' THEN 1
			WHEN '26-35岁' THEN 2
			WHEN '36-45岁' THEN 3
			WHEN '46-55岁' THEN 4
			WHEN '56-65岁' THEN 5
			WHEN '66岁及以上' THEN 6
			ELSE 99
		END
    `

	rows, err := db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var stats []AgeGroup
	for rows.Next() {
		var stat AgeGroup
		err := rows.Scan(&stat.Group, &stat.Count)
		if err != nil {
			return nil, err
		}
		stats = append(stats, stat)
	}

	return stats, rows.Err()
}
