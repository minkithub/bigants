package depot

import (
	"context"

	"github.com/allscape/bigants/grasshopper/app/db/tables"
	"github.com/allscape/bigants/grasshopper/model/iam"
)

// SaveUser 메서드는 iam.UserDepot 인터페이스를 구현한다.
func (a *Adapter) SaveUser(ctx context.Context, ent *iam.User) error {
	var (
		userRow   *tables.UserRow
		userTable = tables.NewUserTable(a.DB)
	)
	userRow = tables.NewUserRow(ent.ID, ent.Created, ent.Password)

	return userTable.Save(ctx, userRow)
}
