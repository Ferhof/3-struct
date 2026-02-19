package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() string {
	c.Key = os.Getenv("KEY")
	return c.Key
}
