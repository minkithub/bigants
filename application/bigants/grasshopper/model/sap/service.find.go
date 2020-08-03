package sap

import (
	"context"
	"fmt"
)

func (s *Service) FindStockOrFail(ctx context.Context, code string) (*Stock, error) {
	stocks, err := s.GetStocksByCode(ctx, code)
	if err != nil {
		return nil, err
	}
	if stocks[0] == nil {
		return nil, fmt.Errorf("Stock: Does not exist(%s)", code)
	}
	return stocks[0], nil
}
