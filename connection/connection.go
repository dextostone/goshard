package connection

import (
	"database/sql"
	"errors"
	"fmt"

	// Importing this package to use mysql driver
	_ "github.com/go-sql-driver/mysql"
	"github.com/vinod4006/goshard/config"
)

func getError(msg string, err error) error {
	if err != nil {
		msg = msg + " " + err.Error()
	}
	return errors.New(msg)
}

// Connection Base type for creating new connections and running queries
type Connection struct {
	db  *sql.DB
	age int
}

// CreateConnection Function to create new connections
func (conn Connection) CreateConnection(dbtype string, configObj config.Config) error {
	if dbtype != "mysql" {
		return getError("dbType must be mysql", nil)
	}

	connectionString := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", configObj.Username,
		configObj.Password, configObj.Hostname, configObj.PortNumber, configObj.Dbname)
	db, err := sql.Open("mysql", connectionString)
	conn.db = db
	if err != nil {
		return getError("Error occured while establishing connection. ", err)
	}
	// Now ping to check if connection is actually established
	err = db.Ping()
	if err != nil {
		return getError("Error occured while establishing connection. ", err)
	}

	return nil

}
