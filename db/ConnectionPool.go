package db

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"sync"
	"time"
)

var (
	db     *sql.DB
	dbOnce sync.Once
)

func GetConnection() (*sql.DB, error) {

	// Performing the creation of ConnectionPool in a singleton function
	dbOnce.Do(func() {
		err := godotenv.Load("db/db.env")
		if err != nil {
			fmt.Println("error loading .env file: %w", err)
			return
		}
		connString1 := os.Getenv("db.url")
		maxOpenConns, err := strconv.Atoi(os.Getenv("db.max"))
		db, err = sql.Open("sqlserver", connString1)
		if err != nil {
			fmt.Println("error creating connection pool: %w", err)
			return
		}
		db.SetMaxOpenConns(maxOpenConns)
		db.SetConnMaxIdleTime(600000 * time.Millisecond)
		db.SetConnMaxLifetime(1200 * 1000 * time.Millisecond)
	})
	return db, nil
}
