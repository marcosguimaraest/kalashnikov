package appdb

import (
	"database/sql"
	"fmt"
)

func AppDatabase() *sql.DB {
	db, err := sql.Open("sqlite3", "./app.db")
	if err != nil {
		return (nil)
	}
	stmt, errprep := db.Prepare("CREATE TABLE IF NOT EXISTS users(number VARCHAR(13) PRIMARY KEY, name VARCHAR(50))")
	stmt.Exec()
	if errprep != nil {
		fmt.Println("\n\n\n ERROR! - ", errprep)
		return (nil)
	}
	stmt, errprep = db.Prepare("INSERT INTO users (number, name) VALUES (\"5521970079469\", \"RONALDO LIMA DOS SANTOS\")")
	if errprep != nil {
		fmt.Println("\n\n\n ERROR! - ", errprep)
		return (nil)
	}
	stmt.Exec()
	return (db)
}
