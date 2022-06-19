package usecase

import (
	"context"
	"time"
	"wiiki_server/domain/model/repomodel"
	"wiiki_server/domain/repository"
)

type User interface {
	Create(ctx context.Context, txTime time.Time, name string, email string, password string) error
	Delete(ctx context.Context, userID string) error
	Update(ctx context.Context, userID string, name *string, email *string, password *string) error
	List(ctx context.Context) ([]*repomodel.User, error)
	Get(ctx context.Context, userID string) (*repomodel.User, error)
}

func NewUser(userRepository repository.User) User {
	return &userImpl{
		userRepository: userRepository,
	}
}

type userImpl struct {
	userRepository repository.User
}

func (impl *userImpl) Create(ctx context.Context, txTime time.Time, name string, email string, password string) error {
	return nil
}

func (impl *userImpl) Delete(ctx context.Context, userID string) error {
	return nil
}

func (impl *userImpl) Update(ctx context.Context, userID string, name *string, email *string, password *string) error {
	return nil
}

func (impl *userImpl) List(ctx context.Context) ([]*repomodel.User, error) {
	return nil, nil
}

func (impl *userImpl) Get(ctx context.Context, userID string) (*repomodel.User, error) {
	return nil, nil
}
