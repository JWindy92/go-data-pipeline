package config

type Config struct {
	App         AppConfig
	Environment EnvironmentConfig
}

type AppConfig struct {
	Name string
}

type EnvironmentConfig struct {
	Production bool
	SeedDb     bool
}
