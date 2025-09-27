package vault

import (
	"demo/files/account"
	"encoding/json"
	"strings"
	"time"

	"github.com/fatih/color"
)

type Db interface {
	Read() ([]byte, error)
	Write([]byte)
}

type Vault struct {
	Accounts  []account.Account `json:"accounts"`
	UpdatedAt time.Time         `json:"updatedAt"`
}

type VaultWithDb struct {
	Vault
	db Db
}

/*
NewVault создает новый экземпляр хранилища Vault.
Если файл accounts.json существует, метод читает его содержимое и десериализует в структуру Vault.
В случае ошибки или отсутствия файла возвращает пустое хранилище Vault с текущим временем обновления.
*/
func NewVault(db Db) *VaultWithDb {
	file, err := db.Read()

	if err != nil {
		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []account.Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	var vault Vault
	err = json.Unmarshal(file, &vault)

	if err != nil {
		color.Red("Error reading file:", err.Error())

		return &VaultWithDb{
			Vault: Vault{
				Accounts:  []account.Account{},
				UpdatedAt: time.Now(),
			},
			db: db,
		}
	}

	return &VaultWithDb{
		Vault: vault,
		db:    db,
	}
}

func (vault *VaultWithDb) FindAccountsByUrl(url string) []account.Account {
	var accounts []account.Account

	for _, account := range vault.Accounts {
		if strings.Contains(account.Url, url) {
			accounts = append(accounts, account)
		}
	}

	return accounts
}

func (vault *VaultWithDb) DeleteAccountByUrl(url string) bool {
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
	vault.save()

	return isDeleted
}

/*
AddNewAccount добавляет новый аккаунт в хранилище Vault.
Метод помещает переданный аккаунт в массив Accounts и обновляет время обновления хранилища.
*/
func (vault *VaultWithDb) AddNewAccount(account account.Account) {
	vault.Accounts = append(vault.Accounts, account)
	vault.save()
}

func (vault *Vault) ToBytes() ([]byte, error) {
	file, err := json.Marshal(vault)

	if err != nil {
		return nil, err
	}

	return file, nil
}

func (vault *VaultWithDb) save() {
	vault.UpdatedAt = time.Now()
	file, err := vault.Vault.ToBytes()

	if err != nil {
		color.Red("Error converting account to bytes:", err)
		return
	}

	vault.db.Write(file)
}
