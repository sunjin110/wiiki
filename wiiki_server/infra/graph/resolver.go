package graph

import (
	"wiiki_server/domain/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	TodoUsecase usecase.Todo
}
