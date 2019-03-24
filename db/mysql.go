package db

import (
	"database/sql"

	"github.com/abel1105/go_gin_boilerplate/config"
	"github.com/go-sql-driver/mysql"
)

var db *sql.DB

func Init() *sql.DB {
	Config := mysql.NewConfig()
	Config.User = config.GetConfig().GetString("db.username")
	Config.Passwd = config.GetConfig().GetString("db.password")
	Config.Addr = config.GetConfig().GetString("db.host")
	Config.DBName = config.GetConfig().GetString("db.database")

	var err error
	db, err = sql.Open("mysql", Config.FormatDSN())
	if err != nil {
		panic(err)
	}
	return db
}

func GetDB() *sql.DB {
	return db
}