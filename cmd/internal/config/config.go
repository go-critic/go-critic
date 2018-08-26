package config

func NewConfig() *Config {
	return &Config{Checkers: make(map[string]*Checker)}
}

type Config struct {
	Version int
	Checkers map[string]*Checker
}

type Checker struct {
	Type string
	Parameters map[string]interface{}
}