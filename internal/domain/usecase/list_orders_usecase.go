package usecase

import (
	contract "github.com/reangeline/go-clean-arch/internal/domain/contract/repository"
	"github.com/reangeline/go-clean-arch/internal/dto"
)

type ListOrdersUsecase struct {
	orderRepository contract.OrderRepositotyInterface
}

func NewListOrdersUsecase(orderRepository contract.OrderRepositotyInterface) *ListOrdersUsecase {
	return &ListOrdersUsecase{
		orderRepository: orderRepository,
	}
}

func (o *ListOrdersUsecase) Execute() ([]*dto.ListOrdersOutput, error) {
	var orderOutput []*dto.ListOrdersOutput
	orders, err := o.orderRepository.ListOrders()
	if err != nil {
		return nil, err
	}

	for _, order := range orders {
		orderOutput = append(orderOutput, &dto.ListOrdersOutput{
			OrderId:   order.OrderId.String(),
			Name:      order.Name,
			TypeOrder: order.TypeOrder,
		})
	}

	return orderOutput, nil
}
