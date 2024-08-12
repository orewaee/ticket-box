package config

import (
	"errors"
	"flag"
	"github.com/ilyakaznacheev/cleanenv"
	"os"
)

type Config struct {
	RedirectUri  string `yaml:"redirect-uri" env:"REDIRECT_URI"`
	ClientId     string `yaml:"client-id" env:"CLIENT_ID"`
	ClientSecret string `yaml:"client-secret" env:"CLIENT_SECRET"`
}

var config = new(Config)

func Get() *Config {
	return config
}

func Load() error {
	path := getPath()
	if path == "" {
		return errors.New("config path is empty")
	}

	if _, err := os.Stat(path); os.IsNotExist(err) {
		return errors.New("config path does not exist")
	}

	if err := cleanenv.ReadConfig(path, config); err != nil {
		return err
	}

	return nil
}

func getPath() string {
	var path string

	flag.StringVar(&path, "config", "config.yaml", "path to config file")
	flag.Parse()

	if path == "" {
		path = os.Getenv("CONFIG_PATH")
	}

	return path
}
