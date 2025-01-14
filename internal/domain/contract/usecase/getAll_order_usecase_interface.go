package contract

import (
	"github.com/reangeline/go-clean-arch/internal/dto"
)

type GetAllOrderUseCaseInterface interface {
	Execute() ([]*dto.ListOrdersOutput, error)
}
