package core

import (
	"encoding/json"
	"github.com/4thel00z/emaild/public/core"
	"github.com/monzo/typhon"
	"io/ioutil"
)

func sendEmail(req typhon.Request) typhon.Response {
	var (
		message core.Message
	)

	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		response := req.Response(&GenericResponse{
			Message: "",
			Error:   "Error while consuming the Request",
		})
		response.StatusCode = 400
		return response
	}

	err = json.Unmarshal(body, &message)

	if err != nil {
		response := req.Response(&GenericResponse{
			Message: "",
			Error:   "Send payload with fields: account (string), to (string[]), cc (string[]),\n " +
				"bcc (string[]), body (base64) or html (base64) or file(string) and delay (int)!",
		})
		response.StatusCode = 400
		return response
	}

	//TODO: add scheduling of email and lookup for account
	//TODO: add calculating of id of scheduled email

	return req.Response(&GenericResponse{
			Message: "Successfully received email request.",
			Error:   "",
		})
}
