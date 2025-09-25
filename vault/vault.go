package vault

import (
	"demo/files/account"
	"encoding/json"
	"time"
)

type Vault struct {
	Accounts  []account.Account `json:"accounts"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

func NewVault() *Vault {
	return &Vault{
		Accounts:  []account.Account{},
		UpdatedAt: time.Now(),
	}
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
