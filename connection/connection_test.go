package connection

import (
	"testing"

	"github.com/vinod4006/goshard/config"
)

var configObj config.Config = config.Config{
	Username:   "root",
	Password:   "example",
	Hostname:   "localhost",
	Dbname:     "hello",
	PortNumber: 3306,
}

func TestCreateConnectionForValidDbType(t *testing.T) {

	conn := Connection{}
	err := conn.CreateConnection("mysql", configObj)
	if err != nil {
		t.Error("CreateConnection should have passed")
	}
}

func TestCreateConnectionForInvalidDbType(t *testing.T) {
	conn := Connection{}
	err := conn.CreateConnection("sql", configObj)
	if err == nil {
		t.Error("CreateConnection should have returned error for invalid dbType")
	}
}
