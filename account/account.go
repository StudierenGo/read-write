package account

import (
	"demo/files/helpers"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

const PASSWORD_LENGTH = 13

type UserAccount struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (account UserAccount) OutputUserData() string {
	return fmt.Sprintf(
		"Dear user %s, your password is %s and it's reference to %s\n", helpers.CapitalizeWord(account.Login), account.Password, account.Url)
}

func (account *UserAccount) generatePassword(n int) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	account.Password = string(result)
}

func UserAccountConstructor(userLogin, userPassword, userUrl string) (*UserAccount, error) {
	_, err := url.ParseRequestURI(userUrl)

	if err != nil {
		return nil, errors.New("invalid URL")
	}

	if userLogin == "" {
		return nil, errors.New("login cannot be empty")
	}

	acc := &UserAccount{
		Login:     userLogin,
		Password:  userPassword,
		Url:       userUrl,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	if userPassword == "" {
		acc.generatePassword(PASSWORD_LENGTH)
	}

	return acc, nil
}
