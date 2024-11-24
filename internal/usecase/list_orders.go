package usecase

import "github.com/devfullcycle/20-CleanArch/internal/entity"

type OrderListOutputDTO struct {
	ID         string
	Price      float64
	Tax        float64
	FinalPrice float64
}

type ListOrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

func NewListOrderUseCase(
	OrderRepository entity.OrderRepositoryInterface,
) *ListOrdersUseCase {
	return &ListOrdersUseCase{
		OrderRepository: OrderRepository,
	}
}

func (l *ListOrdersUseCase) Execute() ([]OrderListOutputDTO, error) {
	orders, err := l.OrderRepository.List()
	if err != nil {
		return nil, err
	}
	var ordersOutput []OrderListOutputDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrderListOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}
	return ordersOutput, nil
}
