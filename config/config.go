package config

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)
func GetDB()(db *sql.DB,err error){
	dbDriver:="mysql"
	dbName:="goshop"
	dbUser:="root"
	dbPass:=""
	db,err = sql.Open(dbDriver,dbUser+":"+dbPass+"@/"+dbName)
	return
}