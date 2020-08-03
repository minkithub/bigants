package assertion_test

import (
	"errors"
	"testing"

	"github.com/allscape/bigants/grasshopper/pkg/assertion"
)

func TestJSONDeepEqual(t *testing.T) {
	var err error
	k := struct{ Value int }{12}
	l := struct{ Value int }{12}
	m := struct{ Value int }{11}

	if err = assertion.JSONDeepEqual(k, l); err != nil {
		t.Fatal(err)
	}
	if err = assertion.JSONDeepEqual(k, m); !errors.Is(err, assertion.ErrFailed) {
		t.Fatal(err)
	}
}
