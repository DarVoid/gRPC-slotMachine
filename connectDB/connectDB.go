package connectDB

import (
	"github.com/darvoid/gRPC-slotMachine/conSQLServer"
	"gorm.io/gorm"
)

func Connect() (*gorm.DB, error) {
	user := "<user>"
	password := "<password>"
	ip := "<ip>"
	port := "<port>"
	database, err := conSQLServer.ConnectToSQLServer(user, password, ip, port)
	if err != nil {
		return nil, err
	}
	return database, nil
}
