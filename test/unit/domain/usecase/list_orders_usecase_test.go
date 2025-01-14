package usecase_test

import (
	"errors"
	"testing"

	"github.com/reangeline/go-clean-arch/internal/domain/entity"
	"github.com/reangeline/go-clean-arch/internal/domain/usecase"
	"github.com/reangeline/go-clean-arch/internal/dto"
	pkg_entity "github.com/reangeline/go-clean-arch/pkg/entity"
	"github.com/stretchr/testify/assert"
)

func TestListOrdersUsecase_Execute_Success(t *testing.T) {
	// Configuração do mock e do caso de uso
	mockRepo := new(MockOrderRepository)
	usecase := usecase.NewListOrdersUsecase(mockRepo)

	pkg1 := pkg_entity.NewID()
	pkg2 := pkg_entity.NewID()

	// Pedidos simulados e saída esperada
	orders := []*entity.Order{
		{OrderId: pkg1, Name: "Order1", TypeOrder: "Standard"},
		{OrderId: pkg2, Name: "Order2", TypeOrder: "Express"},
	}
	expectedOutput := []*dto.ListOrdersOutput{
		{OrderId: pkg1.String(), Name: "Order1", TypeOrder: "Standard"},
		{OrderId: pkg2.String(), Name: "Order2", TypeOrder: "Express"},
	}

	// Configuração do mock para retornar a lista de pedidos
	mockRepo.On("ListOrders").Return(orders, nil)

	// Executando o teste
	orderOutput, err := usecase.Execute()

	// Verificando o comportamento e os resultados
	assert.NoError(t, err)
	assert.Equal(t, expectedOutput, orderOutput)
	mockRepo.AssertExpectations(t)
}

func TestListOrdersUsecase_Execute_Error(t *testing.T) {
	// Configuração do mock e do caso de uso
	mockRepo := new(MockOrderRepository)
	usecase := usecase.NewListOrdersUsecase(mockRepo)

	// Definindo erro esperado
	expectedError := errors.New("error fetching orders")

	// Configuração do mock para simular erro, retornando nil corretamente tipado
	mockRepo.On("ListOrders").Return(([]*entity.Order)(nil), expectedError)

	// Executando o teste
	orderOutput, err := usecase.Execute()

	// Verificando o comportamento e os resultados
	assert.Error(t, err)
	assert.Nil(t, orderOutput)
	assert.Equal(t, expectedError, err)
	mockRepo.AssertExpectations(t)
}
