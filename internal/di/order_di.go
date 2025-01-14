package di

import (
	"database/sql"

	"github.com/reangeline/go-clean-arch/internal/domain/usecase"
	"github.com/reangeline/go-clean-arch/internal/infra/database/postgres/repository"
	"github.com/reangeline/go-clean-arch/internal/presentation/controller"
)

func InitializeOrder(db *sql.DB) (*controller.OrderController, error) {
	orderRepository := repository.NewOrderRepository(db)
	getAllOrderUseCase := usecase.NewListOrdersUsecase(orderRepository)
	createOrderUseCase := usecase.NewCreateOrderUsecase(orderRepository)

	order := controller.NewOrderController(getAllOrderUseCase, createOrderUseCase)

	return order, nil
}
