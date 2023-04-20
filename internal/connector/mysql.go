package connector

import (
	"context"
	"database/sql"
	"log"
	"sync"

	"github.com/fatram/devcode-golang/config"
	_ "github.com/go-sql-driver/mysql"
)

var (
	mysqlDatabase     *sql.DB
	mysqlDatabaseOnce sync.Once
)

type DB interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryContext(ctx context.Context, query string, args ...interface{}) (*sql.Rows, error)
	Prepare(query string) (*sql.Stmt, error)
	PrepareContext(ctx context.Context, query string) (*sql.Stmt, error)
	Exec(query string, args ...interface{}) (sql.Result, error)
	ExecContext(ctx context.Context, query string, args ...interface{}) (sql.Result, error)
	QueryRowContext(ctx context.Context, query string, args ...interface{}) *sql.Row
}

func LoadMysqlDatabase() *sql.DB {
	mysqlDatabaseOnce.Do(func() {
		var err error
		mysqlDatabase, err = sql.Open(`mysql`, config.Configuration().DatabaseURI)
		if err != nil {
			log.Panicf("cannot connect to MySQL %s", err)
		}
		err = mysqlDatabase.Ping()
		if err != nil {
			log.Panicf("cannot ping to MySQL %s", err)
		}
	})
	return mysqlDatabase
}
