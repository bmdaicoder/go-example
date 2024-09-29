package bridge

import (
	"fmt"
	"reflect"
)

type MySqlConnection struct {
	url string
}

func NewMysqlConnection(url string) *MySqlConnection {
	return &MySqlConnection{url}
}

type MySqlDatabase struct {
	conn *MySqlConnection
}

func NewMySqlDatabase(conn *MySqlConnection) *MySqlDatabase {
	return &MySqlDatabase{conn}
}

func (ms *MySqlDatabase) Create(input map[string]any, result any) error {
	fmt.Printf("Creating %s using MySQL\n", reflect.TypeOf(result))
	return nil
}

func (ms *MySqlDatabase) Find(id string, result any) error {
	fmt.Printf("Fetching %s, with id %s, using MySQL\n", reflect.TypeOf(result), id)
	return nil
}
