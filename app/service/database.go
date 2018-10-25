package service

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"deploy-station/app/env"
	"log"
	"strconv"
)

func Db() *sql.DB {
	ENV, _ := env.AppEnv()

	host := ENV.Mysql.Host
	port := strconv.Itoa(int(ENV.Mysql.Port))
	dbname := ENV.Mysql.Dbname
	user := ENV.Mysql.User
	pass := ENV.Mysql.Pass

	db, err := sql.Open("mysql", user+":"+pass+"@tcp("+host+":"+port+")/"+dbname+"?parseTime=true&collation=utf8mb4_general_ci")

	//defer db.Close()

	if err != nil {
		log.Fatalln(err)
	}
	db.SetMaxIdleConns(20)
	db.SetMaxOpenConns(20)
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return db
}
