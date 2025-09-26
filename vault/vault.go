package vault

import (
	"demo/files/account"
	"demo/files/files"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

const VaultFileName = "accounts.json"

type Vault struct {
	Accounts  []account.Account `json:"accounts"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

/*
NewVault создает новый экземпляр хранилища Vault.
Если файл accounts.json существует, метод читает его содержимое и десериализует в структуру Vault.
В случае ошибки или отсутствия файла возвращает пустое хранилище Vault с текущим временем обновления.
*/
func NewVault() *Vault {
	file, err := files.ReadFile(VaultFileName)

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

func (vault *Vault) FindAccountsByUrl(url string) []account.Account {
	var accounts []account.Account

	for _, account := range vault.Accounts {
		isMatch := strings.Contains(account.Url, url)
		if isMatch {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

/*
AddNewAccount добавляет новый аккаунт в хранилище Vault.
Метод помещает переданный аккаунт в массив Accounts и обновляет время обновления хранилища.
*/
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
