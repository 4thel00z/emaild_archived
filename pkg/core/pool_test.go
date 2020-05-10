package core

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func TestParseAccounts(t *testing.T) {
	raw := `{
		"my-gmail-account": {
			"type": 0,
			"address": "smtp.gmail.com:587",
			"host": "smtp.gmail.com",
			"count": 100,
			"user": "test@gmail.com",
			"password": "supersecret123!"
		}
	}
	`

	expected := map[string]Account{
		"my-gmail-account": {
			Type:     0,
			Address:  "smtp.gmail.com:587",
			Count:    100,
			User:     "test@gmail.com",
			Password: "supersecret123!",
			Host:     "smtp.gmail.com",
		},
	}
	reader := strings.NewReader(raw)
	accounts, err := ParseAccounts(reader)

	assert.Nil(t, err, func() string {
		if err != nil {
			return err.Error()
		}
		return ""
	}())

	assert.Equal(t, expected, accounts, "")
}
