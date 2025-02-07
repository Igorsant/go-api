package components

import "github.com/ardanlabs/conf"

type Config struct {
	Port         string `conf:"default::8080,env:PORT"`
	DatabaseName string `conf:"default:go-api,env:DATABASE_NAME"`
	MongoURI     string `conf:"default:mongodb://root:example@localhost:27017,env:MONGO_URI"`
}

func NewConfig() Config {
	cfg := Config{}
	err := conf.Parse([]string{}, "", &cfg)
	if err != nil {
		panic(err)
	}
	return cfg
}
