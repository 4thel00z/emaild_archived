package core

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMessage_UnmarshalJSON(t *testing.T) {
	raw := `{
    "to": ["unittest@golang.com"],
	"cc": ["foo@bar.com"],
	"bcc": ["blindguy@isight.com"],
	"body": "YmFzZTY0IG1lICE=",
	"html": "PGh0bWw+PGJvZHk+YmFzZTY0IG1lITwvYm9keT48L2h0bWw+",
	"file": "/tmp/somefile.html",
	"delay": 300}`

	var message Message
	err := json.Unmarshal([]byte(raw), &message)

	assert.Nil(t, err, func() string {
		if err != nil {
			return err.Error()
		}
		return ""
	}())

	fmt.Println(message)
}
