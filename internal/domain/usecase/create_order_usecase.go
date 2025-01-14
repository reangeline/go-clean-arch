package usecase

import (
	contract "github.com/reangeline/go-clean-arch/internal/domain/contract/repository"
	"github.com/reangeline/go-clean-arch/internal/domain/entity"
	"github.com/reangeline/go-clean-arch/internal/dto"
)

type CreateOrderUsecase struct {
	orderRepository contract.OrderRepositotyInterface
}

func NewCreateOrderUsecase(orderRepository contract.OrderRepositotyInterface) *CreateOrderUsecase {

	return &CreateOrderUsecase{
		orderRepository,
	}
}

func (o *CreateOrderUsecase) Execute(input *dto.CreateOrderInput) error {
	order, err := entity.NewOrder(input.Name, input.TypeOrder)
	if err != nil {
		return err
	}

	err = o.orderRepository.Save(order)

	return err
}
