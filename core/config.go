// Package core Provides shared functionalities like config, database connection, controller defination
// that can be used by modules directory
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
	JwtSecret      string
	IsMigrate      bool
	IsSwagger      bool
	Env            Stage
}

func (c *Config) SetMigrate(isMigrate bool) *Config {
	c.IsMigrate = isMigrate
	return c
}
func (c *Config) SetSwagger(isSwagger bool) *Config {
	c.IsSwagger = isSwagger
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

		dbURL := os.Getenv("DB_URL")
		dbName := os.Getenv("DB_NAME")
		jwtSecret := os.Getenv("JWT_SECRET")
		internelSecret := os.Getenv("INTERNAL_SECRET")
		appName := os.Getenv("APP_NAME")

		config = &Config{
			AppName:        appName,
			InternalSecret: internelSecret,
			Port:           port,
			Database: DatabaseConfig{
				DBURL:  dbURL,
				DBNAME: dbName,
			},
			JwtSecret: jwtSecret,
			Env:       env,
		}
	}
	return config
}

func ConfigInstance() *Config {
	return config
}
