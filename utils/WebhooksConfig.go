package utils

import (
	"os"

	"github.com/astaxie/beego"
	"gopkg.in/yaml.v2"
)

type WebhooksYaml struct {
	Webhooks []Webhook `yaml:"webhooks"`
	Smtp     Smtp      `yaml:"smtp"`
}

type Webhook struct {
	Name    string   `yaml:name`
	Actions []Action `yaml:actions`
}

type Action struct {
	Event string `yaml:event`
	Items []Item `yaml:items`
}

// event item
type Item struct {
	Branch         string `yaml:"branch"`
	RepositoryName string `yaml:"repository_name"`
	Script         string `yaml:"script"`
	Secret         string `yaml:"secret"`
}

type Smtp struct {
	Username string `yaml:username`
	Password string `yaml:password`
	Host     string `yaml:host`
	Port     int    `yaml:port`
	SSL      bool   `yaml:ssl`
}

// read yaml config
func ReadYaml() (*WebhooksYaml, error) {
	path := beego.AppPath + "/" + beego.AppConfig.String("webhooks")

	conf := &WebhooksYaml{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		_ = yaml.NewDecoder(f).Decode(conf)
	}

	return conf, nil
}

func GetSmtp() Smtp {
	port, err := beego.AppConfig.Int("SmtpPort")
	if err != nil {
		port = 465
	}
	ssl, err := beego.AppConfig.Bool("SmtpSsl")
	if err != nil {
		ssl = true
	}
	return Smtp{
		Username: beego.AppConfig.String("SmtpUsername"),
		Password: beego.AppConfig.String("SmtpPassword"),
		Host: beego.AppConfig.String("SmtpHost"),
		Port: port,
		SSL: ssl,
	}
}

func GetDefaultNotifyEmail() []string {
	return []string{beego.AppConfig.String("DefaultNotifyEmail")}
}
