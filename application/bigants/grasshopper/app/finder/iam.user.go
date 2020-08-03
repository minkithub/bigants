package finder

import (
	"context"

	"github.com/allscape/bigants/grasshopper/model/iam"
)

// GetUsers 메서드는 iam.UserFinder를 구현한다.
func (a *Adapter) GetUsers(ctx context.Context, s *iam.Service, ids ...string) (ents []*iam.User, err error) {
	ents = make([]*iam.User, len(ids))
	keys := make([]interface{}, len(ids))
	for i, id := range ids {
		keys[i] = struct {
			ID string `json:"id"`
		}{id}
	}
	if _, err = a.DB.QueryAndReceive(ctx, func(i int) []interface{} {
		ent := s.NewUser()
		ents[i] = ent
		return []interface{}{
			&ent.ID, &ent.Created, &ent.Password,
		}
	}, `
		WITH __k AS (
			SELECT id, ROW_NUMBER() over() __idx
			FROM json_populate_recordset(null::gh.user, $1)
		)
		SELECT id, created, password FROM gh.user JOIN __k USING (id)
		ORDER BY __idx
	`, keys); err != nil {
		return nil, err
	}

	for i := 0; i < len(ids); i++ {
		if ents[i] == nil {
			break
		} else if ents[i].ID != ids[i] {
			copy(ents[i+1:], ents[i:])
			ents[i] = nil
		}
	}

	return ents, nil
}
