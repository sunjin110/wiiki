package psglrepository

import (
	"context"
	"wiiki_server/common"
	"wiiki_server/common/wiikictx"
	"wiiki_server/common/wiikierr"
	"wiiki_server/domain/model/repomodel"
	"wiiki_server/domain/repository"
)

type userRepoImpl struct {
	tableName string
}

func NewUser() repository.User {
	return &userRepoImpl{
		tableName: "users",
	}
}

func (impl *userRepoImpl) List(ctx context.Context) ([]*repomodel.User, error) {
	db, err := wiikictx.GetDB(ctx)
	if err != nil {
		return nil, err
	}
	var userList []*repomodel.User
	err = db.Table(impl.tableName).Find(&userList)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedFindRepository, "table=%s", impl.tableName)
	}
	return userList, nil
}

func (impl *userRepoImpl) Get(ctx context.Context, userID string) (*repomodel.User, error) {
	db, err := wiikictx.GetDB(ctx)
	if err != nil {
		return nil, err
	}

	user := &repomodel.User{}
	isExists, err := db.Table(impl.tableName).Where("id = ?", userID).Get(user)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedGetRepository, "table=%s, userID=%s", impl.tableName, userID)
	}

	if !isExists {
		return nil, nil
	}
	return user, nil
}
func (impl *userRepoImpl) Insert(ctx context.Context, user *repomodel.User) error {
	db, err := wiikictx.GetDB(ctx)
	if err != nil {
		return err
	}

	_, err = db.Table(impl.tableName).Insert(user)
	if err != nil {
		return wiikierr.Bind(err, wiikierr.FailedInsertRepository, "table=%s, data=%v", impl.tableName, user)
	}
	return nil
}
func (impl *userRepoImpl) Delete(ctx context.Context, userID string) error {
	db, err := wiikictx.GetDB(ctx)
	if err != nil {
		return err
	}
	_, err = db.Table(impl.tableName).Where("id = ?", userID).Delete(&repomodel.User{})
	if err != nil {
		return wiikierr.Bind(err, wiikierr.FailedDeleteRepository, "table=%s, userID=%s", impl.tableName, userID)
	}
	return nil
}

func (impl *userRepoImpl) Update(ctx context.Context, userID string, updateUser *repomodel.UpdateUser) error {
	return nil
}

func (*userRepoImpl) generateUpdateMap(user *repomodel.UpdateUser) map[string]interface{} {
	m := map[string]interface{}{
		"name":       user.Name,
		"email":      user.Email,
		"created_at": user.CreatedAt,
		"updated_at": user.UpdatedAt,
	}
	return common.ExcludeNilFromMap(m)
}
