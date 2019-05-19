package utils

import (
	"os"

	"gopkg.in/yaml.v2"
)

//var bm cache.Cache

type WebhooksYaml struct {
	WebHooks []WebHook `yaml:"webhooks"`
}

// webhook config item
type WebHook struct {
	Branch       string `yaml:"branch"`
	Repository   string `yaml:"repository"`
	Organization string `yaml:"organization"`
	Script       string `yaml:"script"`
	Secret       string `yaml:"secret"`
}

// read yaml config
func ReadYaml(path string) (*WebhooksYaml, error) {
	conf := &WebhooksYaml{}
	if f, err := os.Open(path); err != nil {
		return nil, err
	} else {
		_ = yaml.NewDecoder(f).Decode(conf)
	}

	return conf, nil
}

//// load config to cache
//func LoadArgs(config string, scripts string) {
//	var err error
//	bm, err = cache.NewCache("memory", `{"interval":60}`)
//	if err != nil {
//		logs.Error("init cache err: " + err.Error())
//	}
//
//	_ = bm.Put("config", config, 100*12*30*24*time.Hour)
//	_ = bm.Put("scripts", scripts, 100*12*30*24*time.Hour)
//}
//
//// get cache key
//func GetCacheKey(key string) (value interface{}) {
//	value = bm.Get(key)
//	return
//}
