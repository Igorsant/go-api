package components

import "github.com/ardanlabs/conf"

type Config struct {
	Port string `conf:"default:3000,env:PORT"`
}

func NewConfig() Config {
	cfg := Config{}
	err := conf.Parse([]string{}, "", &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
