package goconf

import (
	"github.com/spf13/viper"
	"log"
)

func New(path string) *viper.Viper {
	c := viper.New()
	c.SetConfigType("yaml")
	c.SetConfigName("config")
	c.AddConfigPath(path)
	if err := c.ReadInConfig(); err != nil {
		log.Fatalf("got an error reading file config, error: %s", err)
	}
	return c
}
