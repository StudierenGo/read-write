package main

import (
	"demo/packages/account"
	"demo/packages/helpers"
	"fmt"

	"github.com/fatih/color"
)

func main() {
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
