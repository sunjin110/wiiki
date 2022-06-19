package usecase

import (
	"context"
	"time"
	"wiiki_server/common/utils/idutil"
	"wiiki_server/common/wiikierr"
	"wiiki_server/domain/model/repomodel"
	"wiiki_server/domain/repository"
	"wiiki_server/domain/service"
)

type User interface {
	Create(ctx context.Context, txTime time.Time, name string, email string, password string) error
	Delete(ctx context.Context, userID string) error
	Update(ctx context.Context, userID string, name *string, email *string, password *string) error
	List(ctx context.Context) ([]*repomodel.User, error)
	Get(ctx context.Context, userID string) (*repomodel.User, error)
}

func NewUser(userRepository repository.User, hashService service.Hash) User {
	return &userImpl{
		userRepository: userRepository,
		hashService:    hashService,
	}
}

type userImpl struct {
	userRepository repository.User
	hashService    service.Hash
}

func (impl *userImpl) Create(ctx context.Context, txTime time.Time, name string, email string, password string) error {

	hash, err := impl.hashService.Generate(password)
	if err != nil {
		return wiikierr.Bind(err, wiikierr.FailedHashPassword, "")
	}

	user := &repomodel.User{
		ID:       idutil.New(),
		Name:     name,
		Email:    email,
		Password: hash,
	}

	return impl.userRepository.Insert(ctx, user)
}

func (impl *userImpl) Delete(ctx context.Context, userID string) error {
	return impl.userRepository.Delete(ctx, userID)
}

func (impl *userImpl) Update(ctx context.Context, userID string, name *string, email *string, password *string) error {

	var hashedPassword *string
	if password != nil {
		hash, err := impl.hashService.Generate(*password)
		if err != nil {
			return wiikierr.Bind(err, wiikierr.FailedHashPassword, "")
		}
		hashedPassword = &hash
	}

	updateUser := &repomodel.UpdateUser{
		Name:     name,
		Email:    email,
		Password: hashedPassword,
	}

	return impl.userRepository.Update(ctx, userID, updateUser)
}

func (impl *userImpl) List(ctx context.Context) ([]*repomodel.User, error) {
	return impl.userRepository.List(ctx)
}

func (impl *userImpl) Get(ctx context.Context, userID string) (*repomodel.User, error) {
	return impl.userRepository.Get(ctx, userID)
}
