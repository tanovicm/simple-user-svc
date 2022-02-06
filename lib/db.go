package lib

import (
	"github.com/Kamva/mgm"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func OpenDB() error {
	err := mgm.SetDefaultConfig(nil, "users", options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		return err
	}

	return nil
}
