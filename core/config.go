package core

import (
	"os"
	"strconv"
)

type DatabaseConfig struct {
	DBURL  string
	DBNAME string
}

type Config struct {
	AppName        string
	InternalSecret string
	Port           int
	Database       DatabaseConfig
	JWT_SECRET     string
	IS_MIGRATE     bool
	IS_SWAGGER     bool
	Env            Stage
}

func (c *Config) SetMigrate(isMigrate bool) *Config {
	c.IS_MIGRATE = isMigrate
	return c
}
func (c *Config) SetSwagger(isSwagger bool) *Config {
	c.IS_SWAGGER = isSwagger
	return c
}

var config *Config

func NewConfig(env Stage) *Config {
	if config == nil {
		LoadEnv(Stage(env))
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			port = 8080
		}

		db_url := os.Getenv("DB_URL")
		db_name := os.Getenv("DB_NAME")
		jwt_secret := os.Getenv("JWT_SECRET")
		internal_secret := os.Getenv("INTERNAL_SECRET")
		app_name := os.Getenv("APP_NAME")

		config = &Config{
			AppName:        app_name,
			InternalSecret: internal_secret,
			Port:           port,
			Database: DatabaseConfig{
				DBURL:  db_url,
				DBNAME: db_name,
			},
			JWT_SECRET: jwt_secret,
			Env:        env,
		}
	}
	return config
}

func ConfigInstance() *Config {
	return config
}
