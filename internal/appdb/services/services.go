package services

import (
	"errors"
	"fmt"
	"mguimara/kalashnikov/internal/appdb"
	"mguimara/kalashnikov/internal/models"
)

func GetAllUsers() ([]models.User, error) {
	fmt.Println("OPEN CONNECTION", appdb.DB.DB.Stats().OpenConnections)
	rows, err := appdb.DB.DB.Query("SELECT number, name FROM users")
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

func GetUser(number string) (models.User, error) {
	stmt := "SELECT * FROM users WHERE number=" + number
	u := models.User{}
	rows, err := appdb.DB.DB.Query(stmt)
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

func InsertUser(u models.User) error {
	stmt := "INSERT INTO users (number, name) VALUES (\"" + u.NumberID + "\", \"" + u.Name + "\")"
	fmt.Println(stmt)
	rows, errn := appdb.DB.DB.Query(stmt)
	if errn != nil {
		fmt.Println("\n\n\n ERROR! - ", errn)
		return (errn)
	}
	if rows.Next() {
		fmt.Println("\n\nADCIONOU\n\n")
	}
	return (errn)
}
