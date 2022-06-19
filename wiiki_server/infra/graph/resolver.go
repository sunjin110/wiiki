package graph

import (
	"wiiki_server/domain/usecase"

	"xorm.io/xorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	PostgresEngine *xorm.Engine
	TodoUsecase    usecase.Todo
	UserUsecase    usecase.User
}
