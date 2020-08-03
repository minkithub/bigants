package ident_test

import (
	"testing"

	"github.com/allscape/bigants/grasshopper/pkg/ident"
)

func TestImageID(t *testing.T) {
	id, err := ident.DecodeIDExpecting(
		"Image",
		"Image:gs://media__test.timeandspace.app/images/1eddfcfd-9303-454c-9da5-40c45e756c5f.png",
	)
	if err != nil {
		t.Fatal(err)
	}
	if id != "gs://media__test.timeandspace.app/images/1eddfcfd-9303-454c-9da5-40c45e756c5f.png" {
		t.Fatalf("Unexpected id: %s", id)
	}
}
