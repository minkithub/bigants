package auth_test

import (
	"testing"

	"github.com/allscape/bigants/grasshopper/app/auth"
)

func TestEncodeBasic(t *testing.T) {
	s := auth.Service{}

	t.Log(s.EncodeBasic("jj", "1234"))
}
