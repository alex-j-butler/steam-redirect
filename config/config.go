package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

// Config is a struct representing the configuration of the application.
type Config struct {
	Address string `yaml:"address"`
	Port    int    `yaml:"port"`

	Database struct {
		Username string `yaml:"username"`
		Password string `yaml:"password"`
		Database string `yaml:"database"`
		Charset  string `yaml:"charset"`
	} `yaml:"database"`
}

// Conf is the global configuration.
var Conf Config

// InitialiseConfiguration loads the global configuration from the file, and
// using defaults if the file doesn't exist.
func InitialiseConfiguration() {
	configuration, err := ioutil.ReadFile("./config.yml")

	if err != nil {
		log.Println("No configuration file found, using default configuration.")

		// Default configuration.
		Conf = Config{}
		Conf.Address = ""
		Conf.Port = 3000
		Conf.Database.Username = "root"
		Conf.Database.Password = ""
		Conf.Database.Database = "steam_redirect"
		Conf.Database.Charset = "utf8"
	}

	err = yaml.Unmarshal(configuration, &Conf)

	if err != nil {
		log.Println("Configuration parse error:", err)
	}
}
