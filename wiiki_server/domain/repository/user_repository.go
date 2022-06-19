package repository

import (
	"context"
	"wiiki_server/domain/model/repomodel"
)

type User interface {
	Get(ctx context.Context, userID string) (*repomodel.User, error)
	List(ctx context.Context) ([]*repomodel.User, error)
	Insert(ctx context.Context, user *repomodel.User) error
	Delete(ctx context.Context, userID string) error
}
