package db

import (
	"fmt"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/m-posluszny/go-ynab/src/config"
)

type DBWrite = sqlx.DB
type DBRead = DBWrite

var dbWrite *DBWrite

var dbRead *DBRead

func InitDbs(readInfo config.DbConf, writeInfo config.DbConf) {
	dbRead = connectDb(readInfo)
	dbWrite = connectDb(writeInfo)
	fmt.Println("Loading schema")
	res, err := sqlx.LoadFile(dbWrite, "./src/db/schema.sql")
	fmt.Println("DONE:", res != nil, "ERROR:", err)
}

func GetDbWrite() *DBWrite {
	return dbWrite
}

func GetDbRead() *DBRead {
	return dbRead
}

func connectDb(info config.DbConf) *sqlx.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.Name)
	pgdb, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	return pgdb
}
