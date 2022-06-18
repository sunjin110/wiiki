package presenter

import (
	"wiiki_server/domain/model/repomodel"
	graphmodel "wiiki_server/infra/graph/model"
)

func Todo(todo *repomodel.Todo) *graphmodel.Todo {
	return &graphmodel.Todo{
		ID:   todo.ID,
		Text: todo.Text,
		Done: todo.Done,
		User: nil,
	}
}

func TodoList(todoList []*repomodel.Todo) []*graphmodel.Todo {
	var list []*graphmodel.Todo
	for _, todo := range todoList {
		list = append(list, Todo(todo))
	}
	return list
}
