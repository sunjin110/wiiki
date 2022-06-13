package migrations

import (
	"database/sql"
	"fmt"
	"wiiki_server/common/wiikierr"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upMigrate, downMigrate)
}

func upMigrate(tx *sql.Tx) error {

	queryList := []string{
		`
			create table if not exists todos (
				id varchar(24) not null,
				text varchar(256) not null,
				done boolean not null,
				created_at timestamp,
				updated_at timestamp,
				primary key (id)
			);
		`,
		`
			create table if not exists users (
				id varchar(24) not null,
				username varchar(256) not null,
				password varchar(256) not null,
				primary key (id)
			);
		`,
		`
			create table if not exists links (
				id varchar(24) not null,
				title varchar(255),
				address varchar(255),
				user_id varchar(24),
				foreign key (user_id) references users(id),
				primary key (id)
			);
		`,
	}

	for _, query := range queryList {
		fmt.Println("exec query", query)
		_, err := tx.Exec(query)
		if err != nil {
			return wiikierr.Bind(err, wiikierr.MigrateFailed, "query is %s", query)
		}
	}
	return nil
}

func downMigrate(tx *sql.Tx) error {
	queryList := []string{
		`drop table if exists links;`,
		`drop table if exists users;`,
		`drop table if exists todos;`,
	}
	for _, query := range queryList {
		fmt.Println("exec query", query)
		_, err := tx.Exec(query)
		if err != nil {
			return wiikierr.Bind(err, wiikierr.MigrateFailed, "query is %s", query)
		}
	}
	return nil
}
