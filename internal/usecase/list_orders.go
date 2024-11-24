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
