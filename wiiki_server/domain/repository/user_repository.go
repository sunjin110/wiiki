package repository

import (
	"context"
	"wiiki_server/domain/model/repomodel"
)

type User interface {
	Get(ctx context.Context, userID string) (*repomodel.User, error)
	GetByEmail(ctx context.Context, email string) (*repomodel.User, error)
	List(ctx context.Context) ([]*repomodel.User, error)
	ListByIDList(ctx context.Context, idList []string) ([]*repomodel.User, error)
	Insert(ctx context.Context, user *repomodel.User) error
	Delete(ctx context.Context, userID string) error
	Update(ctx context.Context, userID string, updateUser *repomodel.UpdateUser) error
}
