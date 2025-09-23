package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode"
)

/*
getUserInput запрашивает у пользователя логин, пароль и URL,
используя функцию promptUserData из пакета helpers.
Возвращает введённые значения как три строки.
*/
func GetUserInput() (string, string, string) {
	userLogin := promptUserData("Enter your login")
	userPassword := promptUserData("Enter your password")
	userUrl := promptUserData("Enter your url (yandex/google/etc)")

	return userLogin, userPassword, userUrl
}

/*
CapitalizeWord принимает строку и возвращает её копию,
где первая буква приведена к верхнему регистру, а остальные — к нижнему.
Если строка пустая, возвращает её без изменений.
*/
func CapitalizeWord(s string) string {
	if s == "" {
		return s
	}

	runes := []rune(strings.ToLower(s))
	runes[0] = unicode.ToUpper(runes[0])
	return string(runes)
}

/*
promptUserData выводит приглашение пользователю, считывает введённую строку с консоли,
удаляет лишние пробелы и возвращает результат. В случае ошибки возвращает пустую строку.
*/
func promptUserData(prompt string) string {
	fmt.Print(prompt + ": ")

	reader := bufio.NewReader(os.Stdin)
	userAnswer, err := reader.ReadString('\n')

	if err != nil {
		fmt.Println("Error reading input:", err)
		return ""
	}

	userAnswer = strings.TrimSpace(userAnswer)
	return userAnswer
}
