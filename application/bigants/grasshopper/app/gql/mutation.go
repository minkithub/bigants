package gql

import (
	"context"

	"github.com/allscape/bigants/grasshopper/app/gql/schema"
)

func (*Root) Echo(ctx context.Context, args schema.MutationEchoArgs) *string {
	return args.Input
}
