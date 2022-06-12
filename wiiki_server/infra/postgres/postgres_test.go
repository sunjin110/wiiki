package postgres_test

import (
	"testing"
	"wiiki_server/infra/postgres"
)

// go test -v -count=1 -timeout 30s -run ^Test$ wiiki_server/infra/postgres

func Test(t *testing.T) {

	t.Run("connect", func(t *testing.T) {

		err := postgres.Connect()
		if err != nil {
			t.Fatal(err)
		}

	})

}
