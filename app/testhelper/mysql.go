package testhelper

import (
	"database/sql"
	"fmt"

	"github.com/volatiletech/sqlboiler/v4/queries"
)

//Cleanup 全テーブルTruncateする
func Cleanup(db *sql.DB) {

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	tables := []string{
		"users",
	}

	sqls := make([]string, 0, len(tables)+2)
	sqls = append(sqls, "set foreign_key_checks = 0;")
	for _, tab := range tables {
		sqls = append(sqls, "truncate table "+tab)
	}
	sqls = append(sqls, "set foreign_key_checks = 1;")

	for _, sql := range sqls {
		_, err := queries.Raw(sql).Exec(tx)
		if err != nil {
			fmt.Println(err.Error())
		}
	}

	tx.Commit()
}
