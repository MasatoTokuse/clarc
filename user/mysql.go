package users

import (
	"database/sql"
	"time"

	"github.com/go-sql-driver/mysql"
	_ "github.com/go-sql-driver/mysql"
)

type MySQLConnectionString string

func NewMySQLDB(constr MySQLConnectionString) (*sql.DB, error) {
	db, err := sql.Open("mysql", string(constr))
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NewTestMySQLConnectionString() MySQLConnectionString {
	cnf := mysql.NewConfig()
	cnf.User = "admin"
	cnf.Passwd = "pass"
	cnf.Net = "tcp"
	cnf.Addr = "localhost:3306"
	cnf.DBName = "di"
	cnf.ParseTime = true
	cnf.Loc = time.Local
	return MySQLConnectionString(cnf.FormatDSN())
}
