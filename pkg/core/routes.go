package core

import (
	"github.com/monzo/typhon"
)

func Router(scheduler *Scheduler) typhon.Router {
	r := typhon.Router{}

	r.POST("/email/send",sendEmail(scheduler))

	return r
}
