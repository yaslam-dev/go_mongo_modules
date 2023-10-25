package core

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct{}

func (D *Database) InitializeDatabase() {
	config := ConfigInstance()
	db_url := config.Database.DBURL
	db_name := config.Database.DBNAME

	err := mgm.SetDefaultConfig(nil, db_name, options.Client().ApplyURI(db_url))
	if err != nil {
		panic(err)
	}
}
