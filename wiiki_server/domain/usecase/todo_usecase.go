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
	Update(ctx context.Context, txTime time.Time, todoID string, text *string, done *bool, userID *string) error
	List(ctx context.Context) ([]*repomodel.Todo, error)
	Get(ctx context.Context, todoID string) (*repomodel.Todo, error)
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
		UserID:    userID,
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

func (impl *todoImpl) Update(ctx context.Context, txTime time.Time, todoID string, text *string, done *bool, userID *string) error {
	updateTodo := &repomodel.UpdateTodo{
		Text:      text,
		Done:      done,
		UserID:    userID,
		UpdatedAt: &txTime,
	}
	return impl.todoRepository.Update(ctx, todoID, updateTodo)
}

func (impl *todoImpl) List(ctx context.Context) ([]*repomodel.Todo, error) {
	return impl.todoRepository.List(ctx)
}

func (impl *todoImpl) Get(ctx context.Context, todoID string) (*repomodel.Todo, error) {
	return impl.todoRepository.Get(ctx, todoID)
}
