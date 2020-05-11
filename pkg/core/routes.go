package core

import (
	"github.com/monzo/typhon"
)

func Router() typhon.Router {
	r := typhon.Router{}

	r.POST("/email/send",sendEmail)

	return r
}
