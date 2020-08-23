package users

import (
	"testing"
)

func TestNewMySQLDB(t *testing.T) {
	db, err := NewMySQLDB(NewTestMySQLConnectionString())
	dieIf(err)
	err = db.Ping()
	dieIf(err)
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
