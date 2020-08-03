package date_test

import (
	"testing"
	"time"

	"github.com/allscape/bigants/grasshopper/pkg/date"
)

func TestDate(t *testing.T) {
	var d date.Date
	var err error
	if d, err = date.Parse("2006-01-12"); err != nil {
		t.Fatal(err)
	}

	tt := time.Time(d)
	t.Log(tt.Zone())
	t.Log(*tt.Location())

	t.Log(d.Year(), d.Month(), d.Day())
}
