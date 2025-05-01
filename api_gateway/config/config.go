package config

type Config struct {
	AuthPath string
	Router   string
}

func NewConfig() (*Config, error) {
	// TODO
	return &Config{}, nil
}
