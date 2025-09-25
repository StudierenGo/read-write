package vault

import (
	"demo/files/account"
	"demo/files/files"
	"encoding/json"
	"time"

	"github.com/fatih/color"
)

type Vault struct {
	Accounts  []account.Account `json:"accounts"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func NewVault() *Vault {
	file, err := files.ReadFile("accounts.json")

	if err != nil {
		return &Vault{
			Accounts:  []account.Account{},
			UpdatedAt: time.Now(),
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("Error reading file:", err.Error())
	}

	return &vault
}

func (vault *Vault) AddNewAccount(account account.Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.UpdatedAt = time.Now()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}
