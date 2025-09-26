package account

import (
	"errors"
	"math/rand"
	"net/url"
	"time"

	"github.com/fatih/color"
)

const PASSWORD_LENGTH = 13

type Account struct {
	Login     string    `json:"login"`
	Password  string    `json:"password"`
	Url       string    `json:"url"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (account *Account) generatePassword(n int) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	account.Password = string(result)
}

func NewAccount(userLogin, userPassword, userUrl string) (*Account, error) {
	_, err := url.ParseRequestURI(userUrl)

	if err != nil {
		return nil, errors.New("invalid URL")
	}

	if userLogin == "" {
		return nil, errors.New("login cannot be empty")
	}

	acc := &Account{
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

func (account *Account) Output() {
	color.Green("---------------------------")
	color.Green("Login: %s", account.Login)
	color.Green("Url: %s", account.Url)
	color.Green("---------------------------")
}
