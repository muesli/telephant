package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Account holds the settings for one account
type Account struct {
	ConsumerKey       string
	ConsumerSecret    string
	AccessToken       string
	AccessTokenSecret string
}

// Config holds chirp's config settings
type Config struct {
	Account []Account
	Style   string
}

const (
	configFile = "chirp.conf"
)

// LoadConfig returns the current config as a Config struct
func LoadConfig() Config {
	_, err := os.Stat(configFile)
	if err != nil {
		SaveConfig(Config{
			Style: "Material",
			Account: []Account{
				{
					ConsumerKey:       "your consumer key",
					ConsumerSecret:    "your consumer secret",
					AccessToken:       "your access token",
					AccessTokenSecret: "your access token secret",
				},
			},
		})
		log.Fatal("Config file is missing, but a template was created for you! Please edit ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal(err)
	}

	return config
}

// SaveConfig stores the current config
func SaveConfig(config Config) {
	f, err := os.Create(configFile)
	if err != nil {
		panic(err)
	}
	if err := toml.NewEncoder(f).Encode(config); err != nil {
		panic(err)
	}
}
