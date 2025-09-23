package account

import (
	"demo/packages/helpers"
	"errors"
	"fmt"
	"math/rand"
	"net/url"
	"time"
)

const PASSWORD_LENGTH = 13

type UserAccount struct {
	login    string
	password string
	url      string
}

type UserAccountWithTimeStamp struct {
	createdAt time.Time
	updatedAt time.Time
	UserAccount
}

func (account UserAccount) OutputUserData() string {
	return fmt.Sprintf(
		"Dear user %s, your password is %s and it's reference to %s\n", helpers.CapitalizeWord(account.login), account.password, account.url)
}

func (account *UserAccount) generatePassword(n int) {
	var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	result := make([]rune, n)

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}

	account.password = string(result)
}

func UserAccountWithTimeStampConstructor(userLogin, userPassword, userUrl string) (*UserAccountWithTimeStamp, error) {
	_, err := url.ParseRequestURI(userUrl)

	if err != nil {
		return nil, errors.New("invalid URL")
	}

	if userLogin == "" {
		return nil, errors.New("login cannot be empty")
	}

	acc := &UserAccountWithTimeStamp{
		createdAt: time.Now(),
		updatedAt: time.Now(),
		UserAccount: UserAccount{
			login:    userLogin,
			url:      userUrl,
			password: userPassword,
		},
	}

	if userPassword == "" {
		acc.generatePassword(PASSWORD_LENGTH)
	}

	return acc, nil
}
