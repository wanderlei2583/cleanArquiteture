package graph

import (
	"context"

	"github.com/devfullcycle/20-CleanArch/internal/infra/graph/model"
	"github.com/devfullcycle/20-CleanArch/internal/usecase"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	CreateOrderUseCase usecase.CreateOrderUseCase
	ListOrdersUseCase  usecase.ListOrdersUseCase
}

type queryResolver struct{ *Resolver }

func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

func (r *queryResolver) ListOrders(
	ctx context.Context,
) ([]*model.Order, error) {
	output, err := r.ListOrdersUseCase.Execute()
	if err != nil {
		return nil, err
	}

	var orders []*model.Order
	for _, order := range output {
		orders = append(orders, &model.Order{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return orders, nil
}
