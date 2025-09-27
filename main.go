package main

import (
	"demo/files/account"
	"demo/files/files"
	"demo/files/helpers"
	"demo/files/output"
	"demo/files/vault"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	color.Blue("--------------------------------------------")
	color.Blue("=== Welcome to the User Account Manager! ===")
	color.Blue("--------------------------------------------")

	existingVault := vault.NewVault(files.NewJsonDb("data.json"))

	for {
		choice := getMenuChoice()

		switch choice {
		case "1":
			createAccount(existingVault)
		case "2":
			findAccount(existingVault)
		case "3":
			deleteAccount(existingVault)
		case "4":
			color.Green("Exiting the program. Goodbye!")
			return
		default:
			output.PrintMessage("Invalid choice. Please try again.")
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

func createAccount(vault *vault.VaultWithDb) {
	userLogin, userPassword, userUrl := helpers.GetUserInput()
	account, err := account.NewAccount(userLogin, userPassword, userUrl)

	if err != nil {
		color.Red("Error creating account:", err)
		return
	}

	vault.AddNewAccount(*account)
}

func findAccount(vault *vault.VaultWithDb) {
	url := helpers.PromptUserData("Enter URL to search")
	accounts := vault.FindAccountsByUrl(url)

	for _, account := range accounts {
		account.Output()
	}

	if len(accounts) == 0 {
		color.Red("-----------------------------------------")
		color.Red("No accounts found for the given URL: %s", url)
		color.Red("-----------------------------------------")
	}
}

func deleteAccount(vault *vault.VaultWithDb) {
	url := helpers.PromptUserData("Enter URL to search")
	isDeleted := vault.DeleteAccountByUrl(url)

	if isDeleted {
		color.Yellow("-----------------------------------------")
		color.Yellow("Account(s) with URL containing '%s' deleted.", url)
		color.Yellow("-----------------------------------------")
	} else {
		color.Red("-----------------------------------------")
		color.Red("No accounts found for the given URL: %s", url)
		color.Red("-----------------------------------------")
	}
}
