package core

import (
	dotEnv "github.com/joho/godotenv"
)

type Stage string

const (
	Development Stage = "dev"
	Production  Stage = "prod"
	Test        Stage = "test"
)

func LoadEnv(env Stage) {
	if env == Development {
		if err := dotEnv.Load(".env.dev"); err != nil {
			panic(err)
		}
	} else if env == Production {
		if err := dotEnv.Load(".env.prod"); err != nil {
			panic(err)
		}
	} else if env == Test {
		if err := dotEnv.Load("../.env.test"); err != nil {
			panic(err)
		}
	} else {
		if err := dotEnv.Load(".env"); err != nil {
			panic(err)
		}
	}
}
