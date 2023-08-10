package main

import (
	"database/sql"

	"github.com/heebin2/bulksql/pkg/bulksql"
)

func main() {
	db, err := sql.Open("mysql", "configure")
	if err != nil {
		panic(err)
	}

	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	defer tx.Rollback()

	bc, err := bulksql.NewBatch("INSERT INTO table_name(att1, att2) VALUES (?, ?) ON ~")
	if err != nil {
		panic(err)
	}

	for i := 0; i < 10; i++ {
		if err := bc.Push(1, 2); err != nil {
			panic(err)
		}
	}

	sql, err := bc.Query()
	if err != nil {
		panic(err)
	}

	if _, err = db.Exec(sql, bc.Datas()...); err != nil {
		panic(err)
	}

	tx.Commit()
}
