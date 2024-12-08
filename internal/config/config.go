package config

type Config struct {
	GetHealthRoute string
	PostExample    string
}

func NewConfig() Config {
	return Config{
		GetHealthRoute: "/health",
		PostExample:    "/example",
	}
}
