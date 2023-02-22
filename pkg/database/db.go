package database

import (
	"context"
	"database/sql"
	"ispick-project-21022023/pkg/model"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type DatabaseOps interface {
	// Auth
	InsertOtp(ctx context.Context, mobile string) error
	ValidateOtp(ctx context.Context, mobile string, otp int) error

	// User
	FindUserByMobile(ctx context.Context, mobile string) (int, error)
	FindUserById(ctx context.Context, id int) (model.UserDetails, error)

	// Home Content
	AppVersion(ctx context.Context) (model.AppVersion, error)
}

type Db struct {
	dbconn *sql.DB
}

func InitDatabase() {
	if os.Getenv("DB_USER") == "" || os.Getenv("DB_HOST") == "" || os.Getenv("DB_PORT") == "" || os.Getenv("DB_NAME") == "" {
		log.Fatal("Database credentials not found!")
	}

	dbconn, err := sql.Open("mysql", os.Getenv("DB_USER")+":@("+os.Getenv("DB_HOST")+":"+os.Getenv("DB_PORT")+")/"+os.Getenv("DB_NAME")+"?parseTime=true")
	if err != nil {
		log.Println(err)
	}
	db = dbconn
}

func GetDb() DatabaseOps {
	return &Db{
		dbconn: db,
	}
}

func (db *Db) FindUserById(ctx context.Context, id int) (model.UserDetails, error) {
	var user model.UserDetails
	statement := "SELECT id, name, slug, email, phone, isMobileVerified, isEmailVerified, isGuestUser, fname, lname, is_approved, is_completed FROM users WHERE id = ? ORDER BY created_at DESC LIMIT 1"
	row := db.dbconn.QueryRow(statement, id)

	err := row.Scan(
		&user.Id,
		&user.Name,
		&user.Slug,
		&user.Email,
		&user.Phone,
		&user.IsMobileVerified,
		&user.IsEmailVerified,
		&user.IsGuestUser,
		&user.Fname,
		&user.Lname,
		&user.Is_approved,
		&user.Is_completed,
	)
	if err != nil {
		return user, err
	}
	return user, nil
}
