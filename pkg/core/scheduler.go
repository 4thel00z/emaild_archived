package core

import (
	"errors"
	"fmt"
	"github.com/jordan-wright/email"
	"io/ioutil"
	"os"
	"strings"
	"time"
)

type Scheduler struct {
	pending     chan Message
	exitChannel chan interface{}
	accounts    map[string]Account
}

func NewScheduler(size int, accounts map[string]Account) *Scheduler {
	return &Scheduler{
		pending:     make(chan Message, size),
		exitChannel: make(chan interface{}, 1),
		accounts:    accounts,
	}
}
func (s *Scheduler) Exit() {
	s.exitChannel <- nil
}

func (s *Scheduler) Schedule(message Message) error {
	if !s.hasAccount(message.Account) {
		return errors.New(fmt.Sprintf("could not find the account: %s", message.Account))
	}
	// TODO: add message validation here
	s.pending <- message
	return nil
}

func (s *Scheduler) hasAccount(account string) bool {
	_, ok := s.accounts[account]
	return ok
}

func (s *Scheduler) Run() {
all:
	for {
	inner:
		select {
		case <-s.exitChannel:
			fmt.Println("Cancelling the scheduler...")
			break all
		case message := <-s.pending:
			{
				account := s.accounts[message.Account]
				e := email.NewEmail()
				e.To = message.To
				e.Cc = message.Cc
				e.Bcc = message.Bcc
				e.From = account.User
				e.ReplyTo = message.ReplyTo
				e.Sender = message.Sender

				html := message.HTML.String()
				if html != "" {
					e.HTML = []byte(html)
				}

				body := message.Body.String()
				if body != "" {
					e.Text = []byte(body)
				}

				f, err := os.Open(message.File.String())
				if err != nil {
					break inner

				}
				defer func() {
					err = f.Close()
					if err != nil {
						//TODO: add error handling
					}
				}()
				content, err := ioutil.ReadAll(f)
				if err != nil {
					//TODO: add error handling
					break inner

				}

				if fileName := message.File.String(); strings.HasSuffix(fileName, ".email") || strings.HasSuffix(fileName, ".txt") {
					e.Text = content
				} else if strings.HasSuffix(fileName, ".html") {
					e.HTML = content
				}
				for _, attachment := range message.Attachments {
					_, err := e.AttachFile(attachment)
					if err != nil {
						//TODO: add error handling
						break inner
					}
				}
				err = account.Sender.Send(e, time.Duration(message.Delay)*time.Second)
				if err != nil {
					//TODO: add error handling
					break inner
				}
			}
		}
	}
}
