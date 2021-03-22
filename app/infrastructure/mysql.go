package infrastructure

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type DBConnectionString string

func NewDB(constr DBConnectionString) (*sql.DB, error) {
	db, err := sql.Open("mysql", string(constr))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewTestMySQLConnectionString() DBConnectionString {
	cnf := mysql.NewConfig()
	cnf.User = "admin"
	cnf.Passwd = "pass"
	cnf.Net = "tcp"
	cnf.Addr = "clarc-mysql:3306"
	cnf.DBName = "clarc"
	cnf.ParseTime = true
	cnf.Loc = time.Local
	return DBConnectionString(cnf.FormatDSN())
}

func NewMockFailedDB(constr DBConnectionString) (*sql.DB, error) {
	db, err := sql.Open("mysql", string(constr))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewMockDBConnectionString() DBConnectionString {
	cnf := mysql.NewConfig()
	cnf.User = "admin"
	cnf.Passwd = "pass"
	cnf.Net = "tcp"
	cnf.Addr = "localhost:3306"
	cnf.DBName = "clarc"
	cnf.ParseTime = true
	cnf.Loc = time.Local
	return DBConnectionString(cnf.FormatDSN())
}
