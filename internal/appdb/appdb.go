package appdb

import (
	"database/sql"
	"errors"
	"fmt"
	"mguimara/kalashnikov/internal/helpers"
	"mguimara/kalashnikov/internal/models"

	"go.mau.fi/whatsmeow/types/events"
)

type AppDB struct {
	DB *sql.DB
}

func NewAppDB() *AppDB {
	return &AppDB{
		DB: AppDatabase(),
	}
}

func (appDB *AppDB) GetAllUsers() ([]models.User, error) {
	fmt.Println("OPEN CONNECTION", appDB.DB.Stats().OpenConnections)
	rows, err := appDB.DB.Query("SELECT number, name FROM users")
	if err != nil {
		fmt.Println("\n\n\n ERROR! - ", err)
		return nil, err
	}
	users := []models.User{}
	for rows.Next() {
		u := models.User{}
		err = rows.Scan(&u.NumberID, &u.Name)
		if err != nil {
			fmt.Println("\n\n\n ERROR! - ", err)
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (appDB *AppDB) GetUser(number string) (models.User, error) {
	stmt := "SELECT * FROM users WHERE number=" + number
	u := models.User{}
	rows, err := appDB.DB.Query(stmt)
	if err != nil {
		fmt.Println("\n\n\n ERROR! - ", err)
		return u, err
	}
	if rows.Next() {
		err = rows.Scan(&u.NumberID, &u.Name)
		if err != nil {
			fmt.Println("\n\n\n ERROR! - ", err)
			return u, err
		}
	} else {
		return u, errors.New("ERROR")
	}
	return u, nil
}

func (appDB *AppDB) InsertUser(u models.User) error {
	stmt := "INSERT INTO users (number, name) VALUES (\"" + u.NumberID + "\", \"" + u.Name + "\")"
	fmt.Println(stmt)
	rows, errn := appDB.DB.Query(stmt)
	if errn != nil {
		fmt.Println("\n\n\n ERROR! - ", errn)
		return (errn)
	}
	if rows.Next() {
		fmt.Println("\n\nADCIONOU\n\n")
	}
	return (errn)
}

func (appDB *AppDB) EventHandler(evt interface{}) {
	switch v := evt.(type) {
	case *events.Message:
		err := appDB.InsertUser(models.User{NumberID: "5521992934590", Name: "REINALDO LEAL"})
		if err != nil {
			fmt.Println("\n\n\nERROR - ", err)
		}
		u, err := appDB.GetUser(v.Info.Sender.User)
		if err == nil && !v.Info.IsGroup {
			helpers.Reply(v, "A equipe do TI está armada com uma kalashnikov e estamos verificando!")
			fmt.Println("\n\n\nUser - ", u)
		} else {
			fmt.Println("UE")
		}
	}
}
