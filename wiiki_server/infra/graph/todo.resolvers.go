package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"
	"wiiki_server/infra/graph/model"
	"wiiki_server/infra/graph/presenter"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	todo, err := r.TodoUsecase.Create(ctx, time.Now(), input.Text, input.UserID)
	if err != nil {
		return nil, err
	}
	return presenter.Todo(todo), nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.TodoID) (bool, error) {
	err := r.TodoUsecase.Delete(ctx, input.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input *model.UpdateTodo) (bool, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Todos(ctx context.Context, done *bool) ([]*model.Todo, error) {
	todoList, err := r.TodoUsecase.List(ctx)
	if err != nil {
		return nil, err
	}
	return presenter.TodoList(todoList), nil
}

func (r *queryResolver) Todo(ctx context.Context, todoID string) (*model.Todo, error) {
	panic(fmt.Errorf("not implemented"))
}
