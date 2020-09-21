package mail

import (
	"fmt"
	"goweb/pkg/log"
)

type logger struct {
}

func (l *logger) SendMail(to, subject, body string, attachs ...*Attach) error {
	log.WithFields(map[string]interface{}{
		"to":      to,
		"subject": subject,
		"body":    body,
		"attachs": fmt.Sprintf("%+v", attachs),
	}).Info("send mail")

	return nil
}

var loggerMail Mail = new(logger)
