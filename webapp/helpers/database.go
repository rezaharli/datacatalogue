package helpers

import (
	"fmt"
	"log"
	"os"

	"github.com/eaciit/clit"

	"git.eaciitapp.com/sebar/dbflex"
	_ "git.eaciitapp.com/sebar/dbflex/drivers/mongodb"
)

var dbConnection dbflex.IConnection

func Database() dbflex.IConnection {
	if dbConnection != nil {
		return dbConnection
	}

	log.Println("-> connecting to database server")

	var err error
	dbConnection, err = NewMongodbConnection()
	if err != nil {
		log.Println("-> failed to connect to the database server.", err.Error())
		os.Exit(0)
	}

	return dbConnection
}

func NewMongodbConnection() (dbflex.IConnection, error) {
	dbConf := clit.Config("default", "database", "").(map[string]interface{})
	dbHost := dbConf["host"]
	dbUsername := dbConf["username"]
	dbPassword := dbConf["password"]
	dbName := dbConf["name"]

	connectionString := "mongodb://"

	if dbHost != "" && dbName != "" && dbUsername != "" && dbPassword != "" {
		connectionString = fmt.Sprintf("%s%s:%s@%s/%s", connectionString, dbUsername, dbPassword, dbHost, dbName)
	} else if dbHost != "" && dbName != "" && dbUsername != "" {
		connectionString = fmt.Sprintf("%s%s@%s/%s", connectionString, dbUsername, dbHost, dbName)
	} else if dbHost != "" && dbName != "" {
		connectionString = fmt.Sprintf("%s%s/%s", connectionString, dbHost, dbName)
	} else {
		return nil, fmt.Errorf("Unable to connect to the database server. Please check the configuration")
	}

	conn, err := dbflex.NewConnectionFromURI(connectionString, nil)
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database server. %s", err.Error())
	}

	err = conn.Connect()
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database server. %s", err.Error())
	}

	return conn, nil
}
