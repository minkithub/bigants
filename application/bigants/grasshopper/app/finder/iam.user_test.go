package finder_test

import (
	"context"
	"testing"

	"github.com/allscape/bigants/grasshopper/app/finder"
	"github.com/allscape/bigants/grasshopper/model/iam"
)

func TestIAMGetUsers(t *testing.T) {
	ctx := context.Background()
	tdb, err := prepareTestDB(ctx)
	if err != nil {
		t.Fatal(err)
	}
	defer tdb.Close()

	testUserIDs := []string{
		"fddc8454-aefc-42ac-a20d-d9723a4c7cb8",
		"fddc8454-aefc-42ac-a20d-aaa23a4c7cb8", // null
		"e83c6c1f-8716-4dca-a788-2e3ae32656c9",
	}

	var userFinder iam.UserFinder = &finder.Adapter{tdb}

	users, err := userFinder.GetUsers(ctx, nil, testUserIDs...)
	if err != nil {
		t.Fatal(err)
	}
	if users[0].ID != testUserIDs[0] {
		t.Fatal()
	}
	if users[1] != nil {
		t.Fatal()
	}
	if users[2].ID != testUserIDs[2] {
		t.Fatal()
	}

}
