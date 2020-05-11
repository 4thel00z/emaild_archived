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
		return "nil"
	}
	return string(*val)
}

func (val *SmartString) String() string {
	if val == nil {
		return "nil"
	}
	return string(*val)
}
