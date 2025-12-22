package config

import (
	"log"
	"os"

	"github.com/naoina/toml"
)

type Config struct {
	Info struct {
		Port    string
		Service string
	} `toml:"info"`

	Exp struct {
		ATExp int64
		RtExp int64
	}

	Cassandra struct {
		Keyspace string
	} `toml:"cassandra"`
}

func NewConfig(path string) *Config {
	c := new(Config)

	f, err := os.Open(path)
	if err != nil {
		log.Panicf("Fail to open config path: %v", err)
	}
	err = toml.NewDecoder(f).Decode(c)
	if err != nil {
		log.Panicf("Fail to decode toml: %v", err)
	}
	return c
}
