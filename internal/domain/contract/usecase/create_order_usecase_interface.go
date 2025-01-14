package contract

import (
	"github.com/reangeline/go-clean-arch/internal/dto"
)

type CreateOrderUseCaseInterface interface {
	Execute(input *dto.CreateOrderInput) error
}
