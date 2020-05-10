package core

import (
	"encoding/json"
	"github.com/jordan-wright/email"
	"io"
	"io/ioutil"
	"log"
	"net/smtp"
	"os"
)

const (
	GOOGLE = iota
	MICROSOFT
)

type Account struct {
	Type        int    `json:"type"`
	Address     string `json:"address"`
	Count       int    `json:"count"`
	User        string `json:"user"`
	Password    string `json:"password"`
	Host        string `json:"host"`
	*email.Pool `json:"_"`
}

func ParseAccountsFromPath(path string) (map[string]Account, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}

	defer func() {
		err := file.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	return ParseAccounts(file)
}

func ParseAccounts(reader io.Reader) (map[string]Account, error) {
	var accs map[string]Account

	content, err := ioutil.ReadAll(reader)

	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(content, &accs)

	if err != nil {
		return nil, err
	}

	for _, acc := range accs {
		pool, err := email.NewPool(
			acc.Address,
			acc.Count,
			smtp.PlainAuth("", acc.User, acc.Password, acc.Host),
		)

		if err != nil {
			return nil, err
		}
		acc.Pool = pool
	}

	return accs, nil

}
