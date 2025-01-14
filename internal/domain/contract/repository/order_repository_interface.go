package contract

import "github.com/reangeline/go-clean-arch/internal/domain/entity"

type OrderRepositotyInterface interface {
	Save(order *entity.Order) error
	ListOrders() ([]*entity.Order, error)
}
