package core

import (
	"github.com/4thel00z/emaild/pkg/core"
)

type Message struct {
	// Which of the accounts registered via the pool module you want to schedule on
	Account string   `json:"account"`
	To      []string `json:"to"`
	CC      []string `json:"cc"`
	BCC     []string `json:"bcc"`
	//base64 encoded
	Body *core.Base64 `json:"body,omitempty"`
	//base64 encoded
	HTML *core.Base64     `json:"html,omitempty"`
	File *core.SmartString `json:"file,omitempty"`
	// delay in seconds from now
	Delay int `json:"delay"`
}
