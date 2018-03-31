package datastore

import(
	"errors"
	"tutorials/backendwebdev/gopherface/models"
)

type Datastore interface{
	CreateUser(user *models.User) error
	GetUser(username string) (*models.User, error)
	Close()
}

const (
	MYSQL = iota
	MONGODB
	REDIS
)

func NewDatastore(datastoreType int, dbConnectionString string) (Datastore, error){
	switch datastoreType{
	case MYSQL:
		return NewSQLDatastore(dbConnectionString)
	case MONGODB:
		return NewMongoDBDatastore(dbConnectionString)
	case REDIS:
		return NewRedisDataStore(dbConnectionString)
	default:
		return nil, error.New("The datastore you specified does not exist!")
	}
}