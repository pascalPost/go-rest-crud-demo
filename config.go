package main

type Config struct {
	dev bool
}

func NewConfig(dev bool) *Config {
	return &Config{
		dev: dev,
	}
}

func (c *Config) Dev() bool {
	return c.dev
}
