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
		if strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *Vault) DeleteAccountByUrl(url string) bool {
	var filtered []account.Account
	isDeleted := false

	// - аккаунты, которые не нужно удалять, сразу добавляются в новый срез и цикл продолжается
	// - аккаунты, которые нужно удалить, просто пропускаются (не добавляются в новый срез)
	for _, account := range vault.Accounts {
		if !strings.Contains(account.Url, url) {
			filtered = append(filtered, account)
			continue
		}
		isDeleted = true
	}

	vault.Accounts = filtered
	vault.writeInFile()

	return isDeleted
}

/*
AddNewAccount добавляет новый аккаунт в хранилище Vault.
Метод помещает переданный аккаунт в массив Accounts и обновляет время обновления хранилища.
*/
func (vault *Vault) AddNewAccount(account account.Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.writeInFile()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vault *Vault) writeInFile() {
	vault.UpdatedAt = time.Now()
	file, err := vault.ToBytes()

	if err != nil {
		color.Red("Error converting account to bytes:", err)
		return
	}

	files.WriteFile(VaultFileName, file)
}
