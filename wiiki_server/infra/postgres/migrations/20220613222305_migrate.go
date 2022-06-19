package migrations

import (
	"database/sql"
	"wiiki_server/common/wiikierr"

	"github.com/pressly/goose/v3"
)

func init() {
	goose.AddMigration(upMigrate, downMigrate)
}

func upMigrate(tx *sql.Tx) error {

	queryList := []string{
		`
			create table if not exists users (
				id varchar(24) not null,
				name varchar(256) not null,
				password varchar(256) not null,
				email varchar(255) not null,
				created_at timestamp,
				updated_at timestamp,
				primary key (id)
			);
		`,
		`
			create table if not exists todos (
				id varchar(24) not null,
				text varchar(256) not null,
				done boolean not null,
				user_id varchar(24),
				created_at timestamp,
				updated_at timestamp,
				foreign key (user_id) references users(id),
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
		`drop table if exists todos;`,
		`drop table if exists users;`,
	}
	for _, query := range queryList {
		_, err := tx.Exec(query)
		if err != nil {
			return wiikierr.Bind(err, wiikierr.MigrateFailed, "query is %s", query)
		}
	}
	return nil
}
