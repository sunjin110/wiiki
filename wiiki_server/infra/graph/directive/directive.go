package directive

import (
	"context"
	"fmt"
	"wiiki_server/common/utils/jsonutil"
	"wiiki_server/infra/graph/model"

	"github.com/99designs/gqlgen/graphql"
)

func HasRole(ctx context.Context, obj interface{}, next graphql.Resolver, role model.Role) (interface{}, error) {

	fmt.Println("RUN has role!!!!")

	fmt.Println("========== hasRoleObj is", jsonutil.MustMarshal(obj))

	if role.String() == model.RoleAdmin.String() {
		return nil, fmt.Errorf("Access denied")
	}

	return next(ctx)
}
