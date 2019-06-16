package utils

import (
	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	conf, err := ReadYaml("")

	msg := gomail.NewMessage()
	msg.SetHeader("From","XD Game" + "<" + conf.Smtp.Username + ">")
	msg.SetHeader("To", mailTo...)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", body)

	if err == nil {
		mailer := gomail.NewDialer(conf.Smtp.Host, conf.Smtp.Port, conf.Smtp.Username, conf.Smtp.Password)
		if err := mailer.DialAndSend(msg); err != nil {
			return err
		}
	}

	return nil
}
