package resolver

import (
	"context"

	"github.com/allscape/bigants/grasshopper/app/gql/schema"

	"github.com/allscape/bigants/grasshopper/model/iam"
	"github.com/allscape/bigants/grasshopper/model/list"
	"github.com/allscape/bigants/grasshopper/model/sap"
	"github.com/allscape/bigants/grasshopper/pkg/ident"
	"github.com/graph-gophers/graphql-go"
)

func (rsv *Resolver) ResolveUser(ctx context.Context, id string) (*User, error) {
	ents, err := rsv.iams.GetUsers(ctx, id)
	if err != nil {
		return nil, err
	}
	if ents[0] == nil {
		return nil, nil
	}
	return rsv.NewUser(ents[0]), nil
}

func (rsv *Resolver) NewUser(src *iam.User) *User { return &User{rsv, src} }

type User struct {
	root *Resolver
	src  *iam.User
}

func (rsv *User) ID() graphql.ID { return ident.EncodeID("User", rsv.src.ID) }

func (rsv *User) Predictions(ctx context.Context, args schema.UserPredictionsArgs) (*StockPredictionConnection, error) {
	conn, err := rsv.root.saps.PaginateStockPredictions(ctx,
		list.PageOption{
			args.First, args.After, args.Last, args.Before,
		},
		sap.PredictionFilter{
			RequesterID: &rsv.src.ID,
		},
	)
	if err != nil {
		return nil, err
	}
	return &StockPredictionConnection{rsv.root, conn}, nil
}

func (rsv *User) RecentStocks() []*Stock {
	panic("unimplemented")
}
