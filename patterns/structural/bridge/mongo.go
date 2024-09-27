package bridge

import (
	"fmt"
	"reflect"
)

type MongodbConnection struct {
	url string
}

func NewMongodbConnection(url string) *MongodbConnection {
	return &MongodbConnection{url}
}

type MongoDatabase struct {
	conn *MongodbConnection
}

func NewMongoDatabase(conn *MongodbConnection) *MongoDatabase {
	return &MongoDatabase{conn: conn}
}

func (md MongoDatabase) Create(input map[string]any, result any) error {
	fmt.Printf("Creating %s using MongoDB\n", reflect.TypeOf(result))
	return nil
}

func (md MongoDatabase) Find(id string, result any) error {
	fmt.Printf("Fetching %s, with id %s, using MongoDB\n", reflect.TypeOf(result), id)
	return nil
}
