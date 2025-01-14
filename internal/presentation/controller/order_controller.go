package controller

import (
	"encoding/json"
	"net/http"

	usecase "github.com/reangeline/go-clean-arch/internal/domain/contract/usecase"
	"github.com/reangeline/go-clean-arch/internal/dto"
)

type Error struct {
	Message string `json:"message"`
}

type OrderController struct {
	getAllOrderUseCase usecase.GetAllOrderUseCaseInterface
	createOrderUseCase usecase.CreateOrderUseCaseInterface
}

func NewOrderController(
	getAllOrderUseCase usecase.GetAllOrderUseCaseInterface,
	createOrderUseCase usecase.CreateOrderUseCaseInterface,
) *OrderController {
	return &OrderController{
		getAllOrderUseCase: getAllOrderUseCase,
		createOrderUseCase: createOrderUseCase,
	}
}

func (o *OrderController) GetAllOrder(w http.ResponseWriter, r *http.Request) {

	var order dto.ListOrdersOutput
	err := json.NewDecoder(r.Body).Decode(&order)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(order)

}

func (co *OrderController) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var order dto.CreateOrderInput
	err := json.NewDecoder(r.Body).Decode(&order)
	defer r.Body.Close()

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		error := Error{Message: err.Error()}
		json.NewEncoder(w).Encode(error)
		return
	}

	err = co.createOrderUseCase.Execute(&order)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
