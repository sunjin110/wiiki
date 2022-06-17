package postgres_test

import (
	"testing"
	"wiiki_server/common/testtool"
	"wiiki_server/infra/postgres"
)

// go test -v -count=1 -timeout 30s -run ^Test$ wiiki_server/infra/postgres

func Test(t *testing.T) {

	t.Run("new", func(t *testing.T) {

		conf := testtool.Config()

		engine, err := postgres.New(conf.Postgres[0])
		if err != nil {
			t.Fatal(err)
		}
		err = engine.Ping()
		if err != nil {
			t.Fatal(err)
		}
	})

}
