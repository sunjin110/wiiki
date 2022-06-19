package repository

import (
	"context"
	"wiiki_server/domain/model/repomodel"
)

type Todo interface {
	Get(ctx context.Context, todoID string) (*repomodel.Todo, error)
	List(ctx context.Context) ([]*repomodel.Todo, error)
	Insert(ctx context.Context, todo *repomodel.Todo) error
	Update(ctx context.Context, todoID string, updateTodo *repomodel.UpdateTodo) error
	Delete(ctx context.Context, todoID string) error
}
