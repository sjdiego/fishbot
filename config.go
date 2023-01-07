package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Hostname     string   `yaml:"hostname,omitempty"`
	Port         int      `yaml:"port,omitempty"`
	Nickname     string   `yaml:"nickname,omitempty"`
	Ident        string   `yaml:"ident,omitempty"`
	Realname     string   `yaml:"realname,omitempty"`
	Automodes    string   `yaml:"automodes,omitempty"`
	Channels     []string `yaml:"channels,omitempty"`
	AllowInvites bool     `yaml:"allow_invites,omitempty"`
}

var config Config

func parseConfig() {
	configFile, err := ioutil.ReadFile("config.yml")

	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(configFile, &config)

	if err != nil {
		panic(err)
	}

	log.Printf("Config OK")
}
