package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	db "github.com/hueodev/auth/database"
)

func getInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func Menu() {
	reader := bufio.NewReader(os.Stdin)
	option, _ := getInput("Choose option (L - Login | S - Sign up): ", reader)

	switch option {
	case "l", "L":
		// Checks Username
		var u string

		fmt.Println("Plese enter a valid username and dont use space or uppercase letters")

		// Enter Username
		fmt.Print("Username: ")
		_, err := fmt.Scan(&u)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Enter password
		var p string

		fmt.Print("Password: ")
		_, err = fmt.Scan(&p)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Check credentials
		_, err = db.CheckCredentials(u, p)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

	case "s", "S":
		// Checks Username
		var u string

		fmt.Println("Plese enter a valid username and dont use space or uppercase letters!")

		// Checks Username
		fmt.Print("Username: ")
		_, err := fmt.Scan(&u)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Checks Password
		var p string

		fmt.Print("Password: ")
		_, err = fmt.Scan(&p)
		if err != nil {
			fmt.Println(err.Error())
			return
		}

		// Inserts credentials into DB and checks uf Username is available
		var exist bool
		exist, err = db.CheckUsername(u)
		if err != nil {
			fmt.Println(err.Error())
			return
		}
		if exist {
			fmt.Printf("Username %v is already taken!", u)
		} else {
			_, err = db.InsertDB(u, p)
			if err != nil {
				fmt.Println(err.Error())
				return
			}
		}

	default:
		fmt.Println("That was not a valid option! you dum dum")
		Menu()
	}
}
