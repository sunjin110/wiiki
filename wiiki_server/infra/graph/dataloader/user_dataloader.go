package dataloader

import (
	"context"
	"database/sql"
	"wiiki_server/common/wiikictx"
	"wiiki_server/domain/model/repomodel"

	"github.com/graph-gophers/dataloader/v7"
	"xorm.io/xorm"
)

type UserReader struct {
	conn    *sql.DB
	session *xorm.Session
}

// func (u *UserReader) GetUsers(ctx context.Context, keys dataloader.) []*dataloader.Result {
// return nil
// }

func (u *UserReader) GetUserList(ctx context.Context, keys []string) []*dataloader.Result[*repomodel.User] {

	db, err := wiikictx.GetReadOnlyDB(ctx)

	return nil
}
