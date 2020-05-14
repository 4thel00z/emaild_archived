package core

import (
	"encoding/json"
	"fmt"
	"github.com/monzo/typhon"
	"io/ioutil"
)

func sendEmail(scheduler* Scheduler) typhon.Service {
	return func(req typhon.Request) typhon.Response {
		var (
			message Message
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
				Error: "Send payload with fields: account (string), to (string[]), cc (string[]),\n " +
					"bcc (string[]), body (base64) or html (base64) or file(string) and delay (int)!",
			})
			response.StatusCode = 400
			return response
		}

		hasAccount := scheduler.hasAccount(message.Account)

		if !hasAccount {
			response := req.Response(&GenericResponse{
				Message: "",
				Error:   fmt.Sprintf("account %s not found", message.Account),
			})
			response.StatusCode = 404
			return response
		}

		err = scheduler.Schedule(message)

		if err != nil {
			response := req.Response(&GenericResponse{
				Message: "",
				Error:   "Could not schedule the given message",
			})
			response.StatusCode = 503
			return response
		}

		//TODO: add calculating of id of scheduled email
		return req.Response(&GenericResponse{
			Message: "Successfully received email request.",
			Error:   "",
		})
	}
}
