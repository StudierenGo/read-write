package main

import (
	"demo/files/account"
	"demo/files/files"
	"demo/files/helpers"
	"fmt"
)

func main() {
	createAccount()

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
