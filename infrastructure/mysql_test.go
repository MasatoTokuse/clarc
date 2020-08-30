package infrastructure

import (
	"testing"
)

func TestNewDB(t *testing.T) {
	db, err := NewDB(NewTestMySQLConnectionString())
	dieIf(err)
	err = db.Ping()
	dieIf(err)
}

func dieIf(err error) {
	if err != nil {
		panic(err)
	}
}
