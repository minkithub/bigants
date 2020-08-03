// Package gql 은 GraphQL 쿼리, 뮤테이션, 타입 리졸버를 구현한다.
package gql

import (
	"context"

	"github.com/allscape/bigants/grasshopper/app/gql/resolver"
	"github.com/allscape/bigants/grasshopper/app/gql/schema"
	"github.com/allscape/bigants/grasshopper/model/iam"
	"github.com/allscape/bigants/grasshopper/model/sap"
	"github.com/graph-gophers/graphql-go"
)

type Service struct {
	*graphql.Schema
	WithResolver func(ctx context.Context) context.Context
}

type Root struct {
	saps *sap.Service
}

func New(iams *iam.Service, saps *sap.Service) (*Service, error) {
	schema, err := graphql.ParseSchema(
		string(schema.SchemaGraphql),
		&Root{saps},
		graphql.UseFieldResolvers(),
	)

	if err != nil {
		return nil, err
	}
	return &Service{
		schema,
		func(ctx context.Context) context.Context {
			return context.WithValue(ctx, keyResolver{}, resolver.New(iams, saps))
		},
	}, nil
}

type keyResolver struct{}

func getResolver(ctx context.Context) *resolver.Resolver {
	return ctx.Value(keyResolver{}).(*resolver.Resolver)
}
