package helpers

import (
	"errors"
	"log"
	"strings"

	"github.com/eaciit/toolkit"

	"github.com/eaciit/clit"
	_ "github.com/eaciit/gora"
	"github.com/gchaincl/dotsql"

	"git.eaciitapp.com/sebar/dbflex"
	_ "git.eaciitapp.com/sebar/dbflex/drivers/mongodb"
)

var dbConnection dbflex.IConnection

func Database() dbflex.IConnection {
	if dbConnection != nil {
		return dbConnection
	}

	dbConf := clit.Config("default", "database", "").(map[string]interface{})
	which := dbConf["which"].(string)
	log.Println("-> connecting to", which, "database server")

	var fn func(dbConf map[string]interface{}, which string) (dbflex.IConnection, error)
	switch which {
	case "oracle":
		fn = NewOracleConnection
	case "mongodb":
		fn = NewMongodbConnection
	}

	var err error
	dbConnection, err = fn(dbConf, which)
	if err != nil {
		log.Println("-> failed to connect to the", which, "database server.", err.Error())
		// os.Exit(0)
	}

	return dbConnection
}

func NewMongodbConnection(dbConf map[string]interface{}, which string) (dbflex.IConnection, error) {
	dbWhichConf := dbConf[which].(map[string]interface{})

	dbHost := dbWhichConf["host"]
	dbUsername := dbWhichConf["username"]
	dbPassword := dbWhichConf["password"]
	dbName := dbWhichConf["name"]

	connectionString := "mongodb://"

	if dbHost != "" && dbName != "" && dbUsername != "" && dbPassword != "" {
		connectionString = toolkit.Sprintf("%s%s:%s@%s/%s", connectionString, dbUsername, dbPassword, dbHost, dbName)
	} else if dbHost != "" && dbName != "" && dbUsername != "" {
		connectionString = toolkit.Sprintf("%s%s@%s/%s", connectionString, dbUsername, dbHost, dbName)
	} else if dbHost != "" && dbName != "" {
		connectionString = toolkit.Sprintf("%s%s/%s", connectionString, dbHost, dbName)
	} else {
		return nil, toolkit.Errorf("Unable to connect to the database server. Please check the configuration")
	}

	conn, err := dbflex.NewConnectionFromURI(connectionString, nil)
	if err != nil {
		return nil, toolkit.Errorf("Unable to connect to the database server. %s", err.Error())
	}

	err = conn.Connect()
	if err != nil {
		return nil, toolkit.Errorf("Unable to connect to the database server. %s", err.Error())
	}

	return conn, nil
}

func NewOracleConnection(dbConf map[string]interface{}, which string) (dbflex.IConnection, error) {
	dbWhichConf := dbConf[which].(map[string]interface{})

	var decryptedPassword string
	var err error
	if dbConf["enableEncryption"] != nil && dbConf["enableEncryption"].(bool) == true {
		decryptedPassword, err = Decrypt(dbWhichConf["password"].(string))
		if err != nil {
			return nil, err
		}
	} else {
		decryptedPassword = dbWhichConf["password"].(string)
	}

	connectionString := toolkit.Sprintf("oracle://%s:%s@%s:%s/%s", dbWhichConf["username"], decryptedPassword, dbWhichConf["host"], dbWhichConf["port"], dbWhichConf["service"])
	conn, err := dbflex.NewConnectionFromURI(connectionString, nil)
	if err != nil {
		return nil, toolkit.Errorf("Unable to connect to the database server. %s", err.Error())
	}

	err = conn.Connect()
	if err != nil {
		return nil, toolkit.Errorf("Unable to connect to the database server. %s", err.Error())
	}

	return conn, nil
}

func TruncateSprintf(str string, args ...interface{}) (string, error) {
	n := strings.Count(str, "%s")
	if n > len(args) {
		return "", errors.New("Unexpected string:" + str)
	}
	return toolkit.Sprintf(str, args[:n]...), nil
}

func BuildQueryFromFile(filePath, queryName string, Colnames []string, args ...interface{}) (string, error) {
	dot, err := dotsql.LoadFromFile(filePath)
	if err != nil {
		return "", err
	}

	raw, err := dot.Raw(queryName)
	if err != nil {
		return "", err
	}

	for key, arg := range args {
		replacedArg := strings.ReplaceAll(toolkit.ToString(arg), "'", "''")
		args[key] = replacedArg
	}

	replaced := strings.ReplaceAll(raw, "%", "%%")
	replaced = strings.ReplaceAll(replaced, "?", "%s")

	return TruncateSprintf(replaced, args...)
}
