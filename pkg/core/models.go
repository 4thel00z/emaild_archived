package core

import (
	"encoding/base64"
	"encoding/json"
)

type Base64 string

type SmartString string

type GenericResponse struct {
	Message interface{} `json:"message"`
	Error   string      `json:"error"`
}

func (val *Base64) UnmarshalJSON(b []byte) error {

	var raw string
	if err := json.Unmarshal(b, &raw); err != nil {
		return err
	}

	decoded, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return err
	}
	*val = Base64(decoded)
	return nil
}

func (val *Base64) String() string {
	if val == nil {
		return ""
	}
	return string(*val)
}

func (val *SmartString) String() string {
	if val == nil {
		return ""
	}
	return string(*val)
}


type Message struct {
	// Which of the accounts registered via the pool module you want to schedule on
	Account     string   `json:"account"`
	To          []string `json:"to"`
	Cc          []string `json:"cc"`
	Bcc         []string `json:"bcc"`
	Subject     string   `json:"subject"`
	ReplyTo     []string `json:"reply_to"`
	Sender      string   `json:"sender"`
	Attachments []string `json:"attachments"`
	//base64 encoded
	Body *Base64 `json:"body,omitempty"`
	//base64 encoded
	HTML *Base64      `json:"html,omitempty"`
	File *SmartString `json:"file,omitempty"`
	// delay in seconds from now
	Delay int `json:"delay"`
}

