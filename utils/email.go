package utils

import (
	"github.com/astaxie/beego/logs"
	"gopkg.in/gomail.v2"
)

func SendMail(mailTo []string, subject string, body string) error {
	conf, err := ReadYaml("")

	if err == nil {
		msg := gomail.NewMessage()
		msg.SetHeader("From", "XD Game"+"<"+conf.Smtp.Username+">")
		msg.SetHeader("To", mailTo...)
		msg.SetHeader("Subject", subject)
		msg.SetBody("text/html", body)

		logs.Info(conf.Smtp.Host)
		mailer := gomail.NewPlainDialer(
			conf.Smtp.Host,
			conf.Smtp.Port,
			conf.Smtp.Username,
			conf.Smtp.Password)
		if err := mailer.DialAndSend(msg); err != nil {
			logs.Warning("Send mail fail:", err.Error())
			return err
		}
	}

	return nil
}
