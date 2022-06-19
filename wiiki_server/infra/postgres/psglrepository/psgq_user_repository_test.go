package psglrepository_test

import (
	"fmt"
	"testing"
	"time"
	"wiiki_server/common/testtool"
	"wiiki_server/common/utils/idutil"
	"wiiki_server/domain/model/repomodel"
	"wiiki_server/infra/postgres/psglrepository"

	"github.com/stretchr/testify/assert"
)

func TestUser(t *testing.T) {

	t.Run("UserRepository", func(t *testing.T) {

		t.Run("Insert", func(t *testing.T) {
			ctx, close := testtool.Context()
			defer close(true)
			txTime := time.Now()

			repo := psglrepository.NewUser()
			userID := idutil.New()
			err := repo.Insert(ctx, &repomodel.User{
				ID:        userID,
				Name:      "test",
				Email:     fmt.Sprintf("%s@test.com", userID),
				Password:  "password",
				CreatedAt: txTime,
				UpdatedAt: txTime,
			})
			assert.Nil(t, err)
		})
	})
}
