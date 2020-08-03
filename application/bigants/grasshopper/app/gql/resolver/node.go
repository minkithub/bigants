package resolver

import (
	"context"
	"fmt"

	"github.com/graph-gophers/graphql-go"
	"golang.org/x/sync/errgroup"

	"github.com/allscape/bigants/grasshopper/pkg/ident"
)

func (rsv *Resolver) ResolveNode(ctx context.Context, globalID graphql.ID) (next *Node, err error) {
	typeName, id, err := ident.DecodeID(globalID)
	if err != nil {
		return nil, err
	}
	next = &Node{}
	switch typeName {
	case "Stock":
		next.src, err = rsv.ResolveStock(ctx, id)
	case "StockPrediction":
		next.src, err = rsv.ResolveStockPrediction(ctx, id)
	case "User":
		next.src, err = rsv.ResolveUser(ctx, id)
	default:
		return nil, fmt.Errorf("Invalid TypeName: %s", typeName)
	}
	return next, err
}

func (rsv *Resolver) ResolveNodes(ctx context.Context, globalIDs []graphql.ID) ([]*Node, error) {
	next := make([]*Node, len(globalIDs))
	eg, ctx := errgroup.WithContext(ctx)
	for i, id := range globalIDs {
		i, id := i, id
		eg.Go(func() (err error) {
			next[i], err = rsv.ResolveNode(ctx, id)
			return err
		})
	}
	return next, eg.Wait()
}

type Node struct {
	root *Resolver
	src  node
}

type node interface {
	ID() graphql.ID
}

func (rsv *Node) ID() graphql.ID { return rsv.src.ID() }

func (rsv *Node) ToStock() (*Stock, bool) {
	c, ok := rsv.src.(*Stock)
	return c, ok
}
func (rsv *Node) ToStockPrediction() (*StockPrediction, bool) {
	c, ok := rsv.src.(*StockPrediction)
	return c, ok
}

func (rsv *Node) ToUser() (*User, bool) {
	c, ok := rsv.src.(*User)
	return c, ok
}
