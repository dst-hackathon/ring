package config

import (
	"github.com/BurntSushi/toml"
	"log"
)

type Config struct {
	Database map[string]database
}

type database struct {
	host string
}

func (db database) Host() string {
	if db.host == "" {
		return "mongodb://localhost"
	}
	return db.host
}

func New() *Config {
	return NewByFile("github.com/dst-hackathon/ring/config.toml")
}

func NewByFile(path string) *Config {
	var config Config
	if _, err := toml.DecodeFile(path, &config); err != nil {
		log.Fatal(err)
	}

	return &config
}
