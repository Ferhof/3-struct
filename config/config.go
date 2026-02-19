package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() {
	c.Key = os.Getenv("KEY")
}
