package postgres

import (
	"database/sql"
	"fmt"
	"wiiki_server/infra/common/config"
	"wiiki_server/infra/common/wiikierr"

	"github.com/go-xorm/xorm"

	_ "github.com/lib/pq"
)

// https://github.com/go-xorm/xorm

func New(conf *config.Postgres) (*xorm.Engine, error) {

	engine, err := xorm.NewEngine("postgres", "")
	if err != nil {
		return nil, wiikierr.Bind(err, "FailedCreateXormEngine", "")
	}

	return engine, nil
}

func Connect() error {
	db, err := sql.Open("postgres", "host=127.0.0.1 port=5433 user=sunjin dbname=wiiki sslmode=disable password=alma")
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println("failed db close")
		}
	}()

	if err != nil {
		fmt.Println(err)
		return err
	}

	// insert
	return nil
}
