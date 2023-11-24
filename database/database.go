package database

import (
	"database/sql"
	"fmt"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

// Creates the DB table
func Database() {
	db, err := sql.Open("sqlite3", "./database/test.sqlite")
	if err != nil {
		fmt.Println("Failed to connect to MySql Database")
		panic(err.Error())
	}

	_, err = db.Exec("CREATE TABLE IF NOT EXISTS users (id INTEGER PRIMARY KEY, name TEXT, password TEXT)")
	if err != nil {
		panic(err.Error())
	}

	defer db.Close()
}

// Inserts Username and Password into DB when signingup
func InsertDB(name string, password string) (string, error) {
	db, _ := sql.Open("sqlite3", "./database/test.sqlite")

	name = strings.ToLower(name)

	_, err := db.Exec("INSERT INTO users (name, password) VALUES (?, ?)", name, password)
	if err != nil {
		fmt.Println("Failed to insert value into the Database")
		panic(err.Error())
	}

	fmt.Printf("Welcome %v You created an Account.\n", name)
	defer db.Close()

	return "", nil
}

// Checks if Username exists
func CheckUsername(name string) (bool, error) {
	db, _ := sql.Open("sqlite3", "./database/test.sqlite")
	defer db.Close()

	var count int
	err := db.QueryRow("SELECT COUNT(*) FROM users WHERE name = ?", name).Scan(&count)
	if err != nil {
		panic(err.Error())
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}

// Checks if Password and Username matches
func CheckCredentials(username, password string) (bool, error) {
	db, err := sql.Open("sqlite3", "./database/test.sqlite")
	if err != nil {
		return false, err
	}
	defer db.Close()

	// Check if the username exists
	var exist bool
	err = db.QueryRow("SELECT EXISTS(SELECT 1 FROM users WHERE name = ?)", username).Scan(&exist)
	if err != nil {
		return false, err
	}

	if !exist {
		return true, nil
	}

	// Check if the password matches the username
	var Password string
	err = db.QueryRow("SELECT password FROM users WHERE name = ?", username).Scan(&Password)
	if err != nil {
		return false, err
	}

	if password == Password {
		fmt.Printf("Welcome %v you are logged in!\n", username)
		return true, nil
	} else {
		fmt.Println("Incorrect Username or Password!")
		return false, nil
	}
}
