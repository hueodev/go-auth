package main

import (
	db "github.com/hueodev/auth/database"
	"github.com/hueodev/auth/input"
)

func main() {
	db.Database()
	input.Menu()
}
