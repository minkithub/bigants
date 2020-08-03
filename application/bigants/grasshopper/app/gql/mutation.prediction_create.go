package gql

import (
	"context"

	"github.com/allscape/bigants/grasshopper/model/iam"
	"github.com/kr/pretty"

	"github.com/graph-gophers/graphql-go"

	"github.com/allscape/bigants/grasshopper/app/gql/resolver"
	"github.com/allscape/bigants/grasshopper/app/gql/schema"
	"github.com/allscape/bigants/grasshopper/model/sap"
	"github.com/allscape/bigants/grasshopper/pkg/date"
	"github.com/allscape/bigants/grasshopper/pkg/ident"
)

func (rsv *Root) PredictionCreate(ctx context.Context, args schema.MutationPredictionCreateArgs) (*PredictionCreatePayload, error) {

	actor := iam.GetActor(ctx)

	stockCode, err := ident.DecodeIDExpecting("Stock", args.Input.StockID)
	if err != nil {
		return nil, err
	}
	holidays := make([]date.Date, len(args.Input.Holidays))
	for i, h := range args.Input.Holidays {
		if holidays[i], err = date.Parse(h); err != nil {
			return nil, err
		}
	}

	stock, err := rsv.saps.FindStockOrFail(ctx, stockCode)
	if err != nil {
		return nil, err
	}

	created, err := rsv.saps.CreatePrediction(ctx, sap.CreatePredictionRequest{
		Stock:     stock,
		Start:     stock.PredictablePeriodStart,
		End:       stock.PredictablePeriodEnd,
		Holidays:  holidays,
		Requester: actor,
	})
	if err != nil {
		return nil, err
	}
	pretty.Log(created)
	if err = rsv.saps.SaveStockPrediction(ctx, created); err != nil {
		return nil, err
	}

	return &PredictionCreatePayload{
		getResolver(ctx),
		args.Input.ClientMutationID,
		created,
	}, nil
}

type PredictionCreatePayload struct {
	root             *resolver.Resolver
	ClientMutationID *graphql.ID
	created          *sap.StockPrediction
}

func (rsv *PredictionCreatePayload) Result() resolver.StockPredictionEdge {
	return rsv.root.ResolveStockPredictionEdge(rsv.created.ID)
}
