package database

import (
	"context"
	"log"
)

func (db *Db) FindUserByMobile(ctx context.Context, mobile string) (int, error) {
	statement := "SELECT id FROM users WHERE phone = ? ORDER BY created_at DESC LIMIT 1"
	log.Println("Search:", mobile)
	row := db.dbconn.QueryRow(statement, mobile)
	var user_id int
	err := row.Scan(
		&user_id,
	)
	if err != nil {
		log.Println(err)
		return 0, err
	}
	return user_id, nil
}
