package core

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct{}

func (D *Database) InitializeDatabase() {
	config := ConfigInstance()
	dbURL := config.Database.DBURL
	dbName := config.Database.DBNAME

	err := mgm.SetDefaultConfig(nil, dbName, options.Client().ApplyURI(dbURL))
	if err != nil {
		panic(err)
	}
}
