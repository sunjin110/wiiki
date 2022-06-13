package postgres_test

import (
	"testing"
	"wiiki_server/common/config"
	"wiiki_server/infra/postgres"
)

// go test -v -count=1 -timeout 30s -run ^Test$ wiiki_server/infra/postgres

func Test(t *testing.T) {

	t.Run("new", func(t *testing.T) {
		conf := &config.Postgres{
			Host:     "127.0.0.1",
			Port:     "5432",
			User:     "wiiki_user",
			DBName:   "wiiki",
			Password: "ZnfZxXY3",
		}

		engine, err := postgres.New(conf)
		if err != nil {
			t.Fatal(err)
		}
		err = engine.Ping()
		if err != nil {
			t.Fatal(err)
		}
	})

}
