package db

import (
	"fmt"

	"github.com/DATA-DOG/go-sqlmock"
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

func GetMockDb() (*sqlx.DB, sqlmock.Sqlmock) {
	mockDB, mock, err := sqlmock.New()
	if err != nil {
		panic(err)
	}
	return sqlx.NewDb(mockDB, "sqlmock"), mock
}

func InitMockDbs() (dbReadMock sqlmock.Sqlmock, dbWriteMock sqlmock.Sqlmock) {
	dbRead, dbReadMock = GetMockDb()
	dbWrite, dbWriteMock = GetMockDb()
	return dbReadMock, dbWriteMock
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
	fmt.Printf(psqlInfo)
	pgdb, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		fmt.Println(err)
	}
	return pgdb
}
