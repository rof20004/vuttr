package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq" // import postgresql driver
)

var database *sql.DB

func init() {
	connection := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", os.Getenv("PGHOST"), os.Getenv("PGPORT"), os.Getenv("PGUSER"), os.Getenv("PGPASSWORD"), os.Getenv("PGDATABASE"))

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	database = db
}

// GetConnection - database connection instance
func GetConnection() *sql.DB {
	return database
}
