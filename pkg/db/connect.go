package db

import (
    "database/sql"
    "os"

    _ "github.com/go-sql-driver/mysql"
    "github.com/joho/godotenv"
)
func Connect() *sql.DB {
	err := godotenv.Load()
	if err != nil {
		panic(err.Error())
	}

	db, err := sql.Open("mysql", os.Getenv("DB_ROLE") + ":" + os.Getenv("DB_PASSWORD") + "@/" + os.Getenv("DB_NAME"))
	if err != nil {
		panic(err.Error())
	}
	return db
}