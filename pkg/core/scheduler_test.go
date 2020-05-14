package core

import (
	"fmt"
	"github.com/jordan-wright/email"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
	"time"
)

type MockSender struct {
	expectedEmail   *email.Email
	expectedTimeout time.Duration
	t               *testing.T
}

func (m MockSender) Send(e *email.Email, timeout time.Duration) (err error) {
	assert.Equal(m.t, m.expectedTimeout, timeout, fmt.Sprintf("expected timeout: %d timeout: %d", m.expectedTimeout, timeout))
	assert.Equal(m.t, m.expectedEmail, e, fmt.Sprintf("expected email: %p email: %p", m.expectedEmail, e))
	return nil
}

func TestScheduler_Schedule(t *testing.T) {
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
	reader := strings.NewReader(raw)
	accounts, err := ParseAccounts(reader)
	assert.Nil(t, err, func() string {
		if err != nil {
			return err.Error()
		}
		return ""
	}())

	exceptedEmail := email.NewEmail()
	exceptedEmail.To = []string{"to@me.com"}
	exceptedEmail.Cc = []string{"cc@me.com"}
	exceptedEmail.Text = []byte("")
	account, ok := accounts["my-gmail-account"]
	assert.True(t, ok, "accounts does not contain: my-gmail-account")

	account.Sender = MockSender{
		t:               t,
		expectedEmail:   exceptedEmail,
		expectedTimeout: 10,
	}
	scheduler := NewScheduler(10, accounts)

	go scheduler.Run()

	body := Base64("")

	message := Message{
		Account: "my-gmail-account",
		To:      []string{"to@me.com"},
		Cc:      []string{"cc@me.com"},
		Body:    &body,
		Delay:   0,
	}

	err = scheduler.Schedule(message)
	assert.Nil(t, err, func() string {
		if err != nil {
			return fmt.Sprintf("scheduler.Schedule threw: %s", err.Error())
		}
		return ""
	}())
	time.Sleep(2 * time.Second)
	scheduler.exitChannel <- nil
}
