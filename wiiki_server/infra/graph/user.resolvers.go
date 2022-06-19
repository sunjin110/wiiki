package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"wiiki_server/common/wiikictx"
	"wiiki_server/infra/graph/model"
	"wiiki_server/infra/graph/presenter"
	"wiiki_server/infra/postgres"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {

	ctx, close, err := postgres.WithReadWriteDB(ctx, r.PostgresEngine)
	defer func() {
		close(err)
		wiikictx.AddError(ctx, err)
	}()

	txTime, err := wiikictx.GetTxTime(ctx)
	if err != nil {
		return nil, err
	}

	user, err := r.UserUsecase.Create(ctx, txTime, input.Name, input.Email, input.Password)
	if err != nil {
		return nil, err
	}

	return presenter.User(user), nil
}

func (r *mutationResolver) DeleteUser(ctx context.Context, input model.UserID) (bool, error) {

	ctx, close, err := postgres.WithReadWriteDB(ctx, r.PostgresEngine)
	defer func() {
		close(err)
		wiikictx.AddError(ctx, err)
	}()

	err = r.UserUsecase.Delete(ctx, input.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (bool, error) {

	ctx, close, err := postgres.WithReadWriteDB(ctx, r.PostgresEngine)
	defer func() {
		close(err)
		wiikictx.AddError(ctx, err)
	}()

	txTime, err := wiikictx.GetTxTime(ctx)
	if err != nil {
		return false, err
	}
	err = r.UserUsecase.Update(ctx, txTime, input.ID, input.Name, input.Email, input.Password)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *queryResolver) Users(ctx context.Context) ([]*model.User, error) {

	ctx = postgres.WithReadDB(ctx, r.PostgresEngine)

	userList, err := r.UserUsecase.List(ctx)
	if err != nil {
		wiikictx.AddError(ctx, err)
		return nil, err
	}
	return presenter.UserList(userList), nil
}

func (r *queryResolver) User(ctx context.Context, id *string, email *string) (*model.User, error) {

	ctx = postgres.WithReadDB(ctx, r.PostgresEngine)

	user, err := r.UserUsecase.FindOne(ctx, id, email)
	if err != nil {
		wiikictx.AddError(ctx, err)
		return nil, err
	}
	return presenter.User(user), nil
}
