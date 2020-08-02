package db

import (
	"fmt"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"                //
	_ "github.com/walf443/go-sql-tracer" //
)

var (
	db  *sqlx.DB
	err error
)

// Init Init
func Init() {
	db, err = sqlx.Open(
		"postgres:trace",
		fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
			os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_DATABASE"),
		),
	)
	if err != nil {
		panic(err)
	}
}

// GetDB GetDB
func GetDB() *sqlx.DB {
	return db
}

// Close Close
func Close() {
	db.Close()
}
