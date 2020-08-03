package db_test

import (
	"context"
	"fmt"
	"reflect"
	"sync"
	"testing"
	"time"

	"github.com/allscape/bigants/grasshopper/app/db"
	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

func getPool(ctx context.Context) *pgxpool.Pool {
	config, err := pgxpool.ParseConfig("user=postgres dbname=pgxb sslmode=disable")
	if err != nil {
		panic(err)
	}
	db, err := pgxpool.ConnectConfig(ctx, config)
	if err != nil {
		panic(err)
	}

	db.Exec(ctx, `DROP TABLE todo`)
	_, err = db.Exec(ctx, `
		CREATE TABLE todo (
			id SERIAL PRIMARY KEY,
			content TEXT NOT NULL,
			done TIMESTAMPTZ
		)
	`)
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(ctx, `
		DO $$
			DECLARE i int = 0;
			BEGIN
				LOOP
					IF i > 2000 THEN EXIT; END IF;
					i = i + 1;
					INSERT INTO todo (content, done) VALUES ('My todo ' || i, NULL);
				END LOOP;
			END
		$$;
	`)
	if err != nil {
		panic(err)
	}
	return db
}

func TestPGXB(t *testing.T) {
	ctx := context.Background()
	pool := getPool(ctx)
	s := &testSender{sender: pool}

	b := db.NewBatcher(ctx, s, 50, 150*time.Microsecond)

	var wg sync.WaitGroup
	for i := 0; i < 250; i++ {
		wg.Add(1)
		args := []interface{}{fmt.Sprintf("My todo %d", i)}

		go func() {
			b.BatchExec(`INSERT INTO todo (content, done) VALUES ($1, null)`, args,
				func(res pgconn.CommandTag, err error) {
					wg.Done()
					if err != nil {
						panic(err)
					}
				},
			)
		}()
	}
	wg.Wait()

	println(s.sendCount)
	if !(3 <= s.sendCount) {
		t.Fatal("Min send count")
	}
	if !(s.sendCount < 120) {
		t.Fatal("Max send count")
	}
	t.Logf("Sent %d batches", s.sendCount)

	for i := 0; i < 250; i++ {
		wg.Add(1)
		id := i + 1
		args := []interface{}{id}
		go func() {
			b.BatchQueryRow(`SELECT id FROM todo WHERE id = $1`, args, func(row pgx.Row, err error) {
				wg.Done()
				if err != nil {
					panic(err)
				}
				var resultID int
				err = row.Scan(&resultID)

				if err != nil {
					t.Fatal(err)
				}
				if !reflect.DeepEqual(resultID, id) {
					t.Fatal()
				}
			})
		}()
	}
	wg.Wait()
	t.Logf("Sent %d batches", s.sendCount)
}

type testSender struct {
	sender    db.BatchSender
	sendCount int
	mu        sync.Mutex
}

func (s *testSender) SendBatch(ctx context.Context, b *pgx.Batch) pgx.BatchResults {
	s.mu.Lock()
	s.sendCount = s.sendCount + 1
	s.mu.Unlock()
	return s.sender.SendBatch(ctx, b)
}
