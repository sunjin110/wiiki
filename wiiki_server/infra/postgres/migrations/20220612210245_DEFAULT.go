package main

import (
	"database/sql"
	"log"
)

// Up is executed when this migration is applied
func Up_20220612210245(txn *sql.Tx) {
	result, err := txn.Exec("create table if exists hello(id integer, name varchar(10));")
	if err != nil {
		log.Println("errir is ", err)
	}
	log.Println("result is ", result)
}

// Down is executed when this migration is rolled back
func Down_20220612210245(txn *sql.Tx) {
	txn.Exec("drop table if exists hello;")
}
