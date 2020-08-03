package db_test

import (
	"context"
	"testing"
	"time"

	"github.com/allscape/bigants/grasshopper/app/db"
)

func TestTestDB(t *testing.T) {
	ctx := context.Background()

	tdb, err := db.NewTDB(ctx, db.DatabaseOptions{})
	if err != nil {
		t.Fatal(err)
	}
	defer tdb.Close()
	if err = tdb.Migrate(ctx, "migrations"); err != nil {
		t.Fatal(err)
	}

	ctx = tdb.WithBatcher(ctx, 30*time.Microsecond, 200)

	var b []byte
	_, err = tdb.QueryAndReceive(ctx,
		func(int) []interface{} { return []interface{}{&b} },
		`SELECT json_agg(tables.*) FROM information_schema.tables WHERE table_schema = 'gh'`,
	)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(string(b))
}
