package usecase

import (
	"context"
	"time"
	"wiiki_server/common/utils/idutil"
	"wiiki_server/domain/model/repomodel"
	"wiiki_server/domain/repository"
)

type Todo interface {
	Create(ctx context.Context, txTime time.Time, text string, userID string) (*repomodel.Todo, error)
	Delete(ctx context.Context, todoID string) error
	Update(ctx context.Context, todoID string, text *string, done *bool, userID *string) error
	List(ctx context.Context) ([]*repomodel.Todo, error)
}

func NewTodo(todoRepository repository.Todo) Todo {
	return &todoImpl{
		todoRepository: todoRepository,
	}
}

type todoImpl struct {
	todoRepository repository.Todo
}

func (impl *todoImpl) Create(ctx context.Context, txTime time.Time, text string, userID string) (*repomodel.Todo, error) {

	repoTodo := &repomodel.Todo{
		ID:        idutil.New(),
		Text:      text,
		Done:      false,
		CreatedAt: txTime,
		UpdatedAt: txTime,
	}

	err := impl.todoRepository.Insert(ctx, repoTodo)
	if err != nil {
		return nil, err
	}

	return repoTodo, nil
}

func (impl *todoImpl) Delete(ctx context.Context, todoID string) error {
	return impl.Delete(ctx, todoID)
}

func (*todoImpl) Update(ctx context.Context, todoID string, text *string, done *bool, userID *string) error {
	return nil
}

func (impl *todoImpl) List(ctx context.Context) ([]*repomodel.Todo, error) {
	return impl.todoRepository.List(ctx)
}
