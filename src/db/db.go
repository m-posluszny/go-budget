package db

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
	"github.com/m-posluszny/go-ynab/src/config"
)

type DBWrite = sql.DB
type DBRead = DBWrite

var dbWrite *DBWrite

var dbRead *DBRead

func InitDbs(readInfo config.DbConf, writeInfo config.DbConf) {
	dbRead = connectDb(readInfo)
	dbWrite = connectDb(writeInfo)
}

func GetDbWrite() *DBWrite {
	return dbWrite
}

func GetDbRead() *DBRead {
	return dbRead
}

func connectDb(info config.DbConf) *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		info.Host, info.Port, info.User, info.Password, info.Name)
	pgdb, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	return pgdb
}
