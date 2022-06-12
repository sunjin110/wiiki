package main

import (
	"database/sql"
	"log"
	"wiiki_server/infra/common/utils/jsonutil"
	"wiiki_server/infra/common/wiikierr"
)

// Up is executed when this migration is applied
func Up_20220612210245(txn *sql.Tx) {
	result, err := txn.Exec("create table if not exists hella(id integer, name varchar(10));")
	if err != nil {
		werr := wiikierr.Bind(err, wiikierr.MigrateFailed, "in default")
		wiikierr.StackTrace(werr)
	}
	log.Println("result is ", jsonutil.MustMarshal(result))
}

// Down is executed when this migration is rolled back
func Down_20220612210245(txn *sql.Tx) {
	txn.Exec("drop table if exists hello;")
}
