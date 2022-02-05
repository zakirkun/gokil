package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config interface {
	Get(key string) string
}

type configTmp struct{}

func (c *configTmp) Get(key string) string {
	return os.Getenv(key)
}

func New(filename ...string) Config {
	err := godotenv.Load(filename...)
	if err != nil {
		panic(err)
	}

	return &configTmp{}
}
