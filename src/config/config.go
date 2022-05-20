package config

import (
	"gopkg.in/ini.v1"
)

type Config interface {
	GetMaxPrintCount() int
}

func NewConfig() Config {
	return &ConfigImpl{}
}

type ConfigImpl struct{}

func (c *ConfigImpl) getConfig() *ini.File {
	cfg, err := ini.Load("./config/config.ini")
	if err != nil {
		panic("config file not found")
	}
	return cfg
}

func (c *ConfigImpl) GetMaxPrintCount() int {
	cfg := c.getConfig()
	maxPrintCount := cfg.Section("Common").Key("max_print_count").MustInt()
	return maxPrintCount
}
