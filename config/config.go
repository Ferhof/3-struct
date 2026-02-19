package config

import "os"

type Config struct {
	Key string
}

func NewConfig() *Config {
	return &Config{}
}

func (c *Config) Load() string {
	key := os.Getenv("KEY")
	if key == "" {
		panic("KEY environment variable not set")
	}
	c.Key = os.Getenv("KEY")
	return c.Key
}
