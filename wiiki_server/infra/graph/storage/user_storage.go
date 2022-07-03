package storage

import (
	"context"
	"wiiki_server/common/wiikierr"
	"wiiki_server/domain/model/repomodel"
	"wiiki_server/domain/repository"

	"github.com/graph-gophers/dataloader/v7"
)

type UserReader interface {
	GetUserList(ctx context.Context, keys []string) []*dataloader.Result[*repomodel.User]
}

func NewUserReader(userRepository repository.User) UserReader {
	return &userReaderImpl{
		userRepository: userRepository,
	}
}

type userReaderImpl struct {
	userRepository repository.User
}

func GetUser(ctx context.Context, userID string) (*repomodel.User, error) {

	loaders := For(ctx)
	thunk := loaders.UserLoader.Load(ctx, userID)
	result, err := thunk()
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedFindRepository, "failed get user. userID=%s", userID)
	}
	return result, nil
}

func (u *userReaderImpl) GetUserList(ctx context.Context, keys []string) []*dataloader.Result[*repomodel.User] {
	userList, err := u.userRepository.ListByIDList(ctx, keys)
	if err != nil {
		// TODO error handling
		panic(err)
	}

	var resultList []*dataloader.Result[*repomodel.User]
	for _, user := range userList {
		result := &dataloader.Result[*repomodel.User]{
			Data:  user,
			Error: nil, // TODO
		}
		resultList = append(resultList, result)
	}

	return resultList
}
