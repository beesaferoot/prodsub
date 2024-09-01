package pkg

type Config struct {
	Port int
}

func NewConfig() *Config {
	return &Config{}
}
