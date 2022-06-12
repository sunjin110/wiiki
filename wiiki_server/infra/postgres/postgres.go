package postgres

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

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
