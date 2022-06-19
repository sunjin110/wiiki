package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"wiiki_server/common/wiikictx"
	"wiiki_server/infra/graph/model"
	"wiiki_server/infra/graph/presenter"
)

func (r *mutationResolver) CreateUser(ctx context.Context, input model.NewUser) (*model.User, error) {
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
	err := r.UserUsecase.Delete(ctx, input.ID)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *mutationResolver) UpdateUser(ctx context.Context, input model.UpdateUser) (bool, error) {
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
	userList, err := r.UserUsecase.List(ctx)
	if err != nil {
		return nil, err
	}
	return presenter.UserList(userList), nil
}

func (r *queryResolver) User(ctx context.Context, id *string, email *string) (*model.User, error) {
	// user, err := r.UserUsecase.Get(ctx, )

	// panic(fmt.Errorf("not implemented"))
	return nil, nil
}
