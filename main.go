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

var menu = map[string]func(*vault.VaultWithDb){
	"1": createAccount,
	"2": findAccountByUrl,
	"3": findAccountByLogin,
	"4": deleteAccount,
}

var menuItems = []string{
	"1. Create a new account",
	"2. Find account by URL",
	"3. Find account by login",
	"4. Delete account",
	"5. Exit",
}

func main() {
	color.Blue("--------------------------------------------")
	color.Blue("=== Welcome to the User Account Manager! ===")
	color.Blue("--------------------------------------------")

	existingVault := vault.NewVault(files.NewJsonDb("data.json"))

	for {
		choice := getMenuChoice[string](menuItems...)

		if choice == "5" {
			color.Blue("Goodbye!")
			break
		}

		menuFn := menu[choice]

		if menuFn == nil {
			output.PrintMessage("Invalid choice. Please try again.")
		}

		menuFn(existingVault)
	}
}

func getMenuChoice[T any](arg ...string) (choice string) {
	for i, item := range arg {
		switch i {
		case len(arg) - 1:
			color.Magenta(item)
		case 0:
			color.Green(item)
		case 1, 2:
			color.Yellow(item)
		case 3:
			color.Red(item)
		default:
			color.Yellow(item)
		}
	}

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

func findAccountByUrl(vault *vault.VaultWithDb) {
	url := helpers.PromptUserData("Enter URL to search")
	accounts := vault.FindAccountsByUrl(url, helpers.CheckUrl)

	helpers.ShowOutputMessage(&accounts, url, "URL")
}

func findAccountByLogin(vault *vault.VaultWithDb) {
	login := helpers.PromptUserData("Enter user login to search")
	accounts := vault.FindAccountsByUrl(login, helpers.CheckLogin)

	helpers.ShowOutputMessage(&accounts, login, "login")
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
