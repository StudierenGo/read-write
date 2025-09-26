package main

import (
	"demo/files/account"
	"demo/files/files"
	"demo/files/helpers"
	vlt "demo/files/vault" // <-- renamed import like vlt to avoid conflict with package name
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Blue("--------------------------------------------")
	color.Blue("=== Welcome to the User Account Manager! ===")
	color.Blue("--------------------------------------------")

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
			color.Green("Exiting the program. Goodbye!")
			return
		default:
			color.Magenta("Invalid choice. Please try again.")
		}
	}
}

func getMenuChoice() (choice string) {
	color.Green("1. Create a new account")
	color.Yellow("2. Find account")
	color.Red("3. Delete account")
	color.Magenta("4. Exit")
	color.Cyan("Enter your choice: ")
	fmt.Scanln(&choice)

	return
}

func createAccount() {
	userLogin, userPassword, userUrl := helpers.GetUserInput()
	account, err := account.NewAccount(userLogin, userPassword, userUrl)

	if err != nil {
		color.Red("Error creating account:", err)
		return
	}

	vault := vlt.NewVault()
	vault.AddNewAccount(*account)
	file, err := vault.ToBytes()

	if err != nil {
		color.Red("Error converting account to bytes:", err)
		return
	}

	files.WriteFile(vlt.VaultFileName, file)
}

func findAccount() {}

func deleteAccount() {}
