package config

type ConfigDatabase struct {
	AppName  string `env:"App" env-default:"TRON"`
	AppEnv   string `env:"App" env-default:"DEV"`
	Port     string `env:"App" env-default:"3000"`
	Host     string `env:"App" env-default:"localhost"`
	LogLevel string `env:"App" env-default:"ERROR"`
}

var Cfg ConfigDatabase