package main

import (
	"database/sql"
	"fmt"
	"wiiki_server/infra/common/wiikierr"
)

// Up is executed when this migration is applied
func Up_20220612210245(txn *sql.Tx) {

	queryList := []string{
		`
			create table if not exists todos (
				id serial not null,
				text varchar(256) not null,
				created_at timestamp,
				updated_at timestamp,
				primary key (id)
			);
		`,
		`
			create table if not exists users (
				id serial not null,
				username varchar(256) not null,
				password varchar(256) not null,
				primary key (id)
			);
		`,
		`
			create table if not exists links (
				id serial not null,
				title varchar(255),
				address varchar(255),
				user_id int,
				foreign key (user_id) references users(id),
				primary key (id)
			);
		`,
	}

	for _, query := range queryList {
		fmt.Println("exec query", query)
		_, err := txn.Exec(query)
		if err != nil {
			werr := wiikierr.Bind(err, wiikierr.MigrateFailed, "query is %s", query)
			wiikierr.StackTrace(werr)
			return
		}
	}

}

// Down is executed when this migration is rolled back
func Down_20220612210245(txn *sql.Tx) {
	queryList := []string{
		`drop table if exists links;`,
		`drop table if exists users;`,
		`drop table if exists todos;`,
	}
	for _, query := range queryList {
		fmt.Println("exec query", query)
		_, err := txn.Exec(query)
		if err != nil {
			werr := wiikierr.Bind(err, wiikierr.MigrateFailed, "query is %s", query)
			wiikierr.StackTrace(werr)
			return
		}
	}
}
