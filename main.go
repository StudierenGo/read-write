package main

import (
	"demo/files/account"
	"demo/files/files"
	"demo/files/helpers"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	files.ReadFile("test.txt")
	files.WriteFile("test.txt", "Hello, World!")
	userLogin, userPassword, userUrl := helpers.GetUserInput()
	account, err := account.UserAccountWithTimeStampConstructor(userLogin, userPassword, userUrl)

	if err != nil {
		fmt.Println("Error creating account:", err)
		return
	}

	result := account.OutputUserData()
	color.Cyan("Account created!")
	color.RGB(255, 128, 0).Println(result)
}
