package presenter

import (
	"wiiki_server/domain/model/repomodel"
	graphmodel "wiiki_server/infra/graph/model"
)

func User(user *repomodel.User) *graphmodel.User {

	if user == nil {
		return nil
	}

	return &graphmodel.User{
		ID:    user.ID,
		Name:  user.Name,
		Email: user.Email,
	}
}

func UserList(userList []*repomodel.User) []*graphmodel.User {
	var list []*graphmodel.User
	for _, user := range userList {
		list = append(list, User(user))
	}
	return list
}
