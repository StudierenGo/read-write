package main

import (
	"demo/files/account"
	"demo/files/files"
	"demo/files/helpers"
	"fmt"
)

func main() {
	fmt.Println("Welcome to the User Account Manager!")

	for {
		choice := getMenuChoice()

		switch choice {
		case "1":
			createAccount()
		case "2":
			findAccount()
		case "3":
			deleteAccount()
		case "4":
			fmt.Println("Exiting the program. Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func getMenuChoice() (choice string) {
	fmt.Println("1. Create a new account")
	fmt.Println("2. Find account")
	fmt.Println("3. Delete account")
	fmt.Println("4. Exit")
	fmt.Print("Enter your choice: ")
	fmt.Scanln(&choice)

	return
}

func createAccount() {
	userLogin, userPassword, userUrl := helpers.GetUserInput()
	account, err := account.UserAccountConstructor(userLogin, userPassword, userUrl)

	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

	file, err := account.ToBytes()

	if err != nil {
		fmt.Println("Error converting account to bytes:", err)
		return
	}

	files.WriteFile("account.json", file)
}

func findAccount() {}

func deleteAccount() {}
