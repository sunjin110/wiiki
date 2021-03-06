package psglrepository

import (
	"context"
	"wiiki_server/common"
	"wiiki_server/common/wiikictx"
	"wiiki_server/common/wiikierr"
	"wiiki_server/domain/model/repomodel"
	"wiiki_server/domain/repository"
)

type todoRepoImpl struct {
	tableName string
}

func NewTodo() repository.Todo {
	return &todoRepoImpl{
		tableName: "todos",
	}
}

func (impl *todoRepoImpl) List(ctx context.Context) ([]*repomodel.Todo, error) {

	db, err := wiikictx.GetReadDB(ctx)
	if err != nil {
		return nil, err
	}

	var todoList []*repomodel.Todo

	err = db.Table(impl.tableName).Find(&todoList)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedFindRepository, "table=%s", impl.tableName)
	}

	return todoList, nil
}

func (impl *todoRepoImpl) Get(ctx context.Context, todoID string) (*repomodel.Todo, error) {

	db, err := wiikictx.GetReadDB(ctx)
	if err != nil {
		return nil, err
	}

	todo := &repomodel.Todo{}
	isExists, err := db.Table(impl.tableName).Where("id = ?", todoID).Get(todo)
	if err != nil {
		return nil, wiikierr.Bind(err, wiikierr.FailedGetRepository, "table=%s, todoID=%s", impl.tableName, todoID)
	}

	if !isExists {
		return nil, nil
	}
	return todo, nil
}

func (impl *todoRepoImpl) Insert(ctx context.Context, todo *repomodel.Todo) error {

	db, err := wiikictx.GetDB(ctx)
	if err != nil {
		return err
	}

	_, err = db.Table(impl.tableName).Insert(todo)
	if err != nil {
		return wiikierr.Bind(err, wiikierr.FailedInsertRepository, "table=%s, data=%v", impl.tableName, todo)
	}

	return nil
}

func (impl *todoRepoImpl) Update(ctx context.Context, todoID string, updateTodo *repomodel.UpdateTodo) error {
	db, err := wiikictx.GetDB(ctx)
	if err != nil {
		return err
	}
	_, err = db.Table(impl.tableName).Where("id = ?", todoID).Update(
		impl.generateUpdateMap(updateTodo),
	)
	if err != nil {
		return wiikierr.Bind(err, wiikierr.FailedUpdateRepository, "table=%s, id=%s, update=%v", impl.tableName, todoID, updateTodo)
	}
	return nil
}

func (impl *todoRepoImpl) Delete(ctx context.Context, todoID string) error {

	db, err := wiikictx.GetDB(ctx)
	if err != nil {
		return err
	}

	_, err = db.Table(impl.tableName).Where("id = ?", todoID).Delete(&repomodel.Todo{})
	if err != nil {
		return wiikierr.Bind(err, wiikierr.FailedDeleteRepository, "table=%s, id=%s", impl.tableName, todoID)
	}

	return nil
}

func (*todoRepoImpl) generateUpdateMap(todo *repomodel.UpdateTodo) map[string]interface{} {
	m := map[string]interface{}{
		"text":       todo.Text,
		"done":       todo.Done,
		"user_id":    todo.UserID,
		"created_at": todo.CreatedAt,
		"updated_at": todo.UpdatedAt,
	}
	return common.ExcludeNilFromMap(m)
}
