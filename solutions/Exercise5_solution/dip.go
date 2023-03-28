package main

import (
	"fmt"
	"log"
)

type Database interface {
	Connect() error
	Query(query string) (result []string, err error)
}

type MySQLDatabase struct {
	host     string
	port     int
	username string
	password string
	database string
}

func (db *MySQLDatabase) Connect() error {
	// implementación para conectar a MySQL
	return nil
}

func (db *MySQLDatabase) Query(query string) (result []string, err error) {
	// implementación para realizar una consulta a MySQL
	return nil, nil
}

func main() {
	db := &MySQLDatabase{
		host:     "localhost",
		port:     3306,
		username: "myuser",
		password: "mypassword",
		database: "mydatabase",
	}

	err := db.Connect()
	if err != nil {
		log.Fatal(err)
	}

	// crear una instancia del objeto que depende de la interfaz Database
	myObject := &MyObject{
		database: db,
	}

	// utilizar el objeto que depende de la interfaz Database
	result, err := myObject.QueryDatabase("SELECT * FROM mytable")
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(result)

	// hacer algo con el resultado
}

type MyObject struct {
	database Database
}

func (o *MyObject) QueryDatabase(query string) (result []string, err error) {
	// utilizar la interfaz Database para realizar una consulta
	result, err = o.database.Query(query)
	return result, err
}
