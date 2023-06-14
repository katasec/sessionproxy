package core

type Config struct {
	Port string
}

func NewConfig() *Config {

	return &Config{
		Port: SPROXY_PORT,
	}
}
