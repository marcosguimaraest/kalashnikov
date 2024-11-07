package appdb

import (
	"database/sql"
)

type AppDB struct {
	DB *sql.DB
}

var DB *AppDB

func InitializeAppDatabase() {
	DB = &AppDB{
		DB: AppDatabase(),
	}
}
