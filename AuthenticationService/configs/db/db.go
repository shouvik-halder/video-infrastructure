package db

import (
	"AuthenticationService/configs"
	"database/sql"
	"fmt"

	"github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func SetupDB(cfg *configs.Config) error {
	dbcfg := mysql.NewConfig()

	dbcfg.User = cfg.DB.DBUSER
	dbcfg.Passwd = cfg.DB.DBPASS
	dbcfg.Addr = cfg.DB.DBADDR
	dbcfg.Net = cfg.DB.DBNET
	dbcfg.DBName = cfg.DB.DBNAME
	dbcfg.ParseTime = true

	db, err := sql.Open("mysql", dbcfg.FormatDSN())
	if err != nil {
		fmt.Println("Issue connecting to DB", err.Error())
		return err
	}

	if err := db.Ping(); err != nil {
		fmt.Println("Error pinging DB", err.Error())
		return err
	}
	DB = db

	fmt.Println("db connection established")
	return nil
}

func GetDB() *sql.DB{
	return DB
}
