package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/fatih/color"
)

/*
getUserInput запрашивает у пользователя логин, пароль и URL,
используя функцию promptUserData из пакета helpers.
Возвращает введённые значения как три строки.
*/
func GetUserInput() (string, string, string) {
	userLogin := PromptUserData("Enter your login")
	userPassword := PromptUserData("Enter your password")
	userUrl := PromptUserData("Enter your url (yandex/google/etc)")

	return userLogin, userPassword, userUrl
}

/*
promptUserData выводит приглашение пользователю, считывает введённую строку с консоли,
удаляет лишние пробелы и возвращает результат. В случае ошибки возвращает пустую строку.
*/
func PromptUserData(prompt string) string {
	fmt.Print(prompt + ": ")

	reader := bufio.NewReader(os.Stdin)
	userAnswer, err := reader.ReadString('\n')

	if err != nil {
		color.Red("Error reading input:", err)
		return ""
	}

	userAnswer = strings.TrimSpace(userAnswer)
	return userAnswer
}
