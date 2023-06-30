package main

import (
	"log"
	"strings"

	"github.com/spf13/viper"
)

type Config struct {
	Hostname       string `mapstructure:"HOSTNAME"`
	Port           int    `mapstructure:"PORT"`
	Nickname       string `mapstructure:"NICKNAME"`
	Ident          string `mapstructure:"IDENT"`
	Realname       string `mapstructure:"REALNAME"`
	Automodes      string `mapstructure:"AUTOMODES"`
	DefaultChannel string `mapstructure:"DEFAULT_CHANNEL"`
	AllowInvites   bool   `mapstructure:"ALLOW_INVITES"`
}

var config Config

func parseConfig() {
	viper.AddConfigPath(".")
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		panic(err)
	}

	err = viper.Unmarshal(&config)

	if err != nil {
		panic(err)
	}

	log.Printf("Config OK")
}
