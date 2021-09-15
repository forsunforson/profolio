package config

import (
	"fmt"
	"os"

	yaml "gopkg.in/yaml.v3"
)

type GlobalConfig struct {
	DateSource []DateSource `yaml:"data_source"`
}

type DateSource struct {
	Name   string `yaml:"name"`
	AppKey string `yaml:"app_key"`
}

var (
	globalConfig *GlobalConfig
)

func InitGlobalConfig() {
	initGlobalConfig()
}

func GetGlobalConfig() *GlobalConfig {
	return globalConfig
}

func initGlobalConfig() {
	f, err := os.Open("./config/config.yaml")
	if err != nil {
		fmt.Printf("cannot open config: %v", err.Error())
	}
	defer f.Close()
	buff := make([]byte, 2046)
	l, err := f.Read(buff)
	if err != nil {
		fmt.Printf("read config fail: %v", err.Error())
	}
	c := GlobalConfig{}
	err = yaml.Unmarshal(buff[:l], &c)
	if err != nil {
		fmt.Printf("unmarshal fail: %v", err.Error())
	}
	globalConfig = &c
}
