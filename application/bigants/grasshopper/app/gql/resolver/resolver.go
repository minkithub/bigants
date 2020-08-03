// Package resolver 는 GraphQL 스키마의 각 타입별 리졸버를 구현한다.
package resolver

import (
	"github.com/allscape/bigants/grasshopper/model/iam"
	"github.com/allscape/bigants/grasshopper/model/sap"
)

func New(iams *iam.Service, saps *sap.Service) *Resolver {
	return &Resolver{iams, saps}
}

type Resolver struct {
	iams *iam.Service
	saps *sap.Service
}
