package db

import (
	"context"
	"database/sql"
	"fmt"
	"log/slog"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/m-posluszny/go-ynab/src/config"
)

type DBWrite = sqlx.DB
type DBRead = DBWrite

var dbWrite *DBWrite

var dbRead *DBRead

// WithTransaction creates a new transaction and handles rollback/commit based on the
// error object returned by the `TxFn`
func Transact(fn func(dbx *sqlx.Tx) error) error {
	tx, err := dbWrite.Beginx()
	if err != nil {
		slog.Error("Transaction Begin Error", "err", err)
		return err
	}

	err = fn(tx)
	if err != nil {
		slog.Error("Transaction Fn Error", "err", err)
		tx.Rollback()
		return nil
	}

	return tx.Commit()
}

type Queryable interface {
	sqlx.ExecerContext
	sqlx.PreparerContext
	sqlx.QueryerContext
	sqlx.Preparer
	sqlx.Execer
	sqlx.Ext

	GetContext(context.Context, interface{}, string, ...interface{}) error
	SelectContext(context.Context, interface{}, string, ...interface{}) error
	Get(interface{}, string, ...interface{}) error
	MustExecContext(context.Context, string, ...interface{}) sql.Result
	PreparexContext(context.Context, string) (*sqlx.Stmt, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
	Select(interface{}, string, ...interface{}) error
	QueryRow(string, ...interface{}) *sql.Row
	PrepareNamedContext(context.Context, string) (*sqlx.NamedStmt, error)
	PrepareNamed(string) (*sqlx.NamedStmt, error)
	Preparex(string) (*sqlx.Stmt, error)
	NamedExec(string, interface{}) (sql.Result, error)
	NamedExecContext(context.Context, string, interface{}) (sql.Result, error)
	MustExec(string, ...interface{}) sql.Result
	NamedQuery(string, interface{}) (*sqlx.Rows, error)
}

func InitDbs(readInfo config.DbConf, writeInfo config.DbConf) {
	slog.Info("Connecting to read db")
	dbRead = connectDb(readInfo)
	slog.Info("Connecting to write db")
	dbWrite = connectDb(writeInfo)
	slog.Info("Loading schema")
	res, err := sqlx.LoadFile(dbWrite, "./src/db/schema.sql")
	if err != nil {
		panic(err)
	}
	slog.Info("DONE:", "response", res != nil, "ERROR:", err)

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
	slog.Info("Connecting to db")
	slog.Info("DB Conn params")
	slog.Info(psqlInfo)
	pgdb, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		slog.Error(err.Error())
	}
	return pgdb
}
