package utils

import (
	"github.com/astaxie/beego"
	"os"

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
	Branch       string `yaml:"branch"`
	Repository   string `yaml:"repository"`
	Organization string `yaml:"organization"`
	Script       string `yaml:"script"`
	Secret       string `yaml:"secret"`
}

type Smtp struct {
	Username     string `yaml:username`
	Password string `yaml:password`
	Host     string `yaml:host`
	Port     int    `yaml:port`
	SSL      bool   `yaml:ssl`
}

// read yaml config
func ReadYaml(path string) (*WebhooksYaml, error) {
	if path == "" {
		path = beego.AppPath + "/" + beego.AppConfig.String("webhooks")
	}

	conf := &WebhooksYaml{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		_ = yaml.NewDecoder(f).Decode(conf)
	}

	return conf, nil
}
