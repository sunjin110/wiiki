package psglrepository_test

import (
	"testing"
	"time"
	"wiiki_server/common/ptr"
	"wiiki_server/common/testtool"
	"wiiki_server/common/utils/idutil"
	"wiiki_server/domain/model/repomodel"
	"wiiki_server/infra/postgres/psglrepository"

	"github.com/stretchr/testify/assert"
)

// go test -v -count=1 -timeout 30s -run ^TestTodo$ wiiki_server/infra/postgres/psglrepository

func TestTodo(t *testing.T) {

	t.Run("TodoRepository", func(t *testing.T) {
		t.Run("Insert", func(t *testing.T) {
			ctx, close := testtool.Context()
			defer close(true)
			txTime := time.Now()

			// test user
			userRepo := psglrepository.NewUser()
			userID := idutil.New()
			err := userRepo.Insert(ctx, &repomodel.User{
				ID:        userID,
				Name:      "test_user",
				Email:     "test@test.com",
				Password:  "password",
				CreatedAt: txTime,
				UpdatedAt: txTime,
			})
			if !assert.Nil(t, err) {
				panic(err)
			}

			repo := psglrepository.NewTodo()
			err = repo.Insert(ctx, &repomodel.Todo{
				ID:        idutil.New(),
				Text:      "sunjin",
				Done:      true,
				UserID:    userID,
				CreatedAt: txTime,
				UpdatedAt: txTime,
			})
			assert.Nil(t, err)
		})

		t.Run("Find", func(t *testing.T) {
			ctx, close := testtool.Context()
			defer close(true)

			repo := psglrepository.NewTodo()
			list, err := repo.List(ctx)
			assert.Nil(t, err)
			assert.NotNil(t, list)

			todoID := list[0].ID

			t.Run("Get", func(t *testing.T) {
				todo, err := repo.Get(ctx, todoID)
				assert.Nil(t, err)
				assert.NotNil(t, todo)
				assert.Equal(t, todo.ID, todoID)

				t.Run("Update", func(t *testing.T) {

					updateTodo := &repomodel.UpdateTodo{
						Done: ptr.ToPtr(true),
						Text: ptr.ToPtr("helloooooo!!"),
					}

					err := repo.Update(ctx, todoID, updateTodo)
					assert.Nil(t, err)

					updatedTodo, err := repo.Get(ctx, todoID)
					assert.Nil(t, err)
					assert.NotNil(t, updatedTodo)
					assert.Equal(t, updatedTodo.Done, true)
					assert.Equal(t, updatedTodo.Text, "helloooooo!!")

					t.Run("Delete", func(t *testing.T) {
						err := repo.Delete(ctx, todoID)
						assert.Nil(t, err)
					})
				})

			})

		})

	})

}
