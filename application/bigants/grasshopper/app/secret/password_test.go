package secret_test

import (
	"testing"

	"github.com/allscape/bigants/grasshopper/app/secret"
)

func TestPasswordHashing(t *testing.T) {
	hashed := "argon2$argon2i$v=19$m=512,t=2,p=2$dFA2eGxlenVsMURx$xsOl+dpeUcXV7o9yvwqi5Q" // from django

	res, err := secret.ComparePasswordAndHash("mystrongpassword", hashed)
	if err != nil {
		t.Fatal(err)
	}

	if res != true {
		t.Fatal("Assertion failed: the given hash was exptected to match mystrongpassword")
	}

	res, err = secret.ComparePasswordAndHash("mystrongpassword2", hashed)
	if err != nil {
		t.Fatal(err)
	}

	if res != false {
		t.Fatal("Assertion failed: the given hash was exptected not to match mystrongpassword2")
	}

	newPassword, err := secret.CreateHashFromPassword("poporipopori")
	if err != nil {
		t.Fatal(err)
	}
	res, err = secret.ComparePasswordAndHash("poporipopori", newPassword)
	if err != nil {
		t.Fatal(err)
	}
	if res != true {
		t.Fatal("Assertion failed: the given hash was expected to match poporipopori")
	}
}

func TestGeneratingRandomPassword(t *testing.T) {
	a := secret.Adapter{}
	plain, err := a.RandomPassword()
	if err != nil {
		t.Fatal(err)
	}
	if len(plain) != 44 {
		t.Fatal()
	}
}
