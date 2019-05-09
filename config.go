package main

import (
	"log"
	"os"

	"github.com/BurntSushi/toml"
)

// Account holds the settings for one account
type Account struct {
	/*
		ConsumerKey       string
		ConsumerSecret    string
		AccessToken       string
		AccessTokenSecret string
	*/
	Instance     string
	ClientID     string
	ClientSecret string
	Token        string
	RedirectURI  string
}

// Config holds telephant's config settings
type Config struct {
	Account []Account
	Style   string
}

// LoadConfig returns the current config as a Config struct
func LoadConfig(configFile string) Config {
	_, err := os.Stat(configFile)
	if err != nil {
		SaveConfig(configFile, Config{
			Style:   "Material",
			Account: []Account{Account{}},
		})
		//log.Fatal("Config file is missing, but a template was created for you! Please edit ", configFile)
	}

	var config Config
	if _, err := toml.DecodeFile(configFile, &config); err != nil {
		log.Fatal("Could not decode config file: ", err)
	}

	return config
}

// SaveConfig stores the current config
func SaveConfig(configFile string, config Config) {
	f, err := os.Create(configFile)
	if err != nil {
		log.Fatal("Could not open config file: ", err)
	}
	if err := toml.NewEncoder(f).Encode(config); err != nil {
		log.Fatal("Could not encode config: ", err)
	}
}
