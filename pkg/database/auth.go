package database

import (
	"context"
	"errors"
	"log"
	"math/rand"
	"time"
)

func (db *Db) InsertOtp(ctx context.Context, mobile string) error {
	rand.Seed(time.Now().UnixNano())
	min := 1000
	max := 9999
	otp := rand.Intn(max-min+1) + min
	query := "insert into otps (mobile, otp_hash, created_at, updated_at) values(?,?,NOW(),NOW())"
	_, err := db.dbconn.ExecContext(ctx, query, mobile, otp)
	return err
}

func (db *Db) ValidateOtp(ctx context.Context, mobile string, otp int) error {
	log.Println(mobile, otp)
	statement := "SELECT id, mobile, otp_hash FROM otps where mobile = ? AND otp_hash = ? AND NOW() < created_at + interval 10 minute ORDER BY 1 desc"
	row := db.dbconn.QueryRow(statement, mobile, otp)

	var (
		id         string
		usermobile string
		otphash    string
	)
	err := row.Scan(&id, &usermobile, &otphash)

	if err != nil {
		log.Println("database:", err)
		return errors.New(err.Error())
	}

	return nil
}
