package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"wiiki_server/common/wiikictx"
	"wiiki_server/infra/graph/generated"
	"wiiki_server/infra/graph/model"
	"wiiki_server/infra/graph/presenter"
	"wiiki_server/infra/graph/storage"
	"wiiki_server/infra/postgres"
)

func (r *mutationResolver) CreateTodo(ctx context.Context, input model.NewTodo) (*model.Todo, error) {
	ctx, close, err := postgres.WithReadWriteDB(ctx, r.PostgresEngine)
	defer func() {
		close(err)
		wiikictx.AddError(ctx, err)
	}()

	txTime, err := wiikictx.GetTxTime(ctx)
	if err != nil {
		return nil, err
	}

	todo, err := r.TodoUsecase.Create(ctx, txTime, input.Text, input.UserID)
	if err != nil {
		return nil, err
	}
	return presenter.Todo(todo), nil
}

func (r *mutationResolver) DeleteTodo(ctx context.Context, input model.TodoID) (bool, error) {
	ctx, close, err := postgres.WithReadWriteDB(ctx, r.PostgresEngine)
	defer func() {
		close(err)
		wiikictx.AddError(ctx, err)
	}()

	err = r.TodoUsecase.Delete(ctx, input.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateTodo(ctx context.Context, input *model.UpdateTodo) (bool, error) {
	ctx, close, err := postgres.WithReadWriteDB(ctx, r.PostgresEngine)
	defer func() {
		close(err)
		wiikictx.AddError(ctx, err)
	}()

	if err != nil {
		return false, err
	}

	txTime, err := wiikictx.GetTxTime(ctx)
	if err != nil {
		return false, err
	}
	err = r.TodoUsecase.Update(ctx, txTime, input.ID, input.Text, input.Done, input.UserID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Todos(ctx context.Context, done *bool) ([]*model.Todo, error) {
	ctx = postgres.WithReadDB(ctx, r.PostgresEngine)
	todoList, err := r.TodoUsecase.List(ctx)
	defer func() {
		wiikictx.AddError(ctx, err)
	}()
	if err != nil {
		return nil, err
	}
	return presenter.TodoList(todoList), nil
}

func (r *queryResolver) Todo(ctx context.Context, todoID string) (*model.Todo, error) {
	ctx = postgres.WithReadDB(ctx, r.PostgresEngine)

	todo, err := r.TodoUsecase.Get(ctx, todoID)
	defer func() {
		wiikictx.AddError(ctx, err)
	}()

	if err != nil {
		return nil, err
	}
	return presenter.Todo(todo), nil
}

func (r *todoResolver) User(ctx context.Context, obj *model.Todo) (*model.User, error) {
	ctx = postgres.WithReadDB(ctx, r.PostgresEngine)
	user, err := storage.GetUser(ctx, obj.UserID)
	defer func() {
		wiikictx.AddError(ctx, err)
	}()
	return presenter.User(user), nil
}

// Todo returns generated.TodoResolver implementation.
func (r *Resolver) Todo() generated.TodoResolver { return &todoResolver{r} }

type todoResolver struct{ *Resolver }
