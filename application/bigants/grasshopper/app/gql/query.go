package gql

import (
	"context"

	"github.com/allscape/bigants/grasshopper/app/gql/resolver"
	"github.com/allscape/bigants/grasshopper/app/gql/schema"
	"github.com/allscape/bigants/grasshopper/model/iam"
)

func (*Root) Version() string {
	return "0.1"
}

func (*Root) ModelUsageCount() int32 {
	return 0
}

func (*Root) Node(ctx context.Context, args schema.QueryNodeArgs) (*resolver.Node, error) {
	if args.ID == nil {
		return nil, nil
	}
	return getResolver(ctx).ResolveNode(ctx, *args.ID)
}

func (*Root) Viewer(ctx context.Context) *resolver.User {
	user := iam.GetActor(ctx)
	if user == nil {
		return nil
	}
	return getResolver(ctx).NewUser(user)
}
