package usecase_test

import (
	"errors"
	"testing"

	"github.com/reangeline/go-clean-arch/internal/domain/entity"
	"github.com/reangeline/go-clean-arch/internal/domain/usecase"
	"github.com/reangeline/go-clean-arch/internal/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Mock do repositório

// Mock do repositório
type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Save(order *entity.Order) error {
	args := m.Called(order)
	return args.Error(0)
}

func (m *MockOrderRepository) ListOrders() ([]*entity.Order, error) {
	args := m.Called()
	return args.Get(0).([]*entity.Order), args.Error(1)
}

func TestCreateOrderUsecase_Execute_Success(t *testing.T) {
	// Configuração do mock e do caso de uso
	mockRepo := new(MockOrderRepository)
	usecaseCreate := usecase.NewCreateOrderUsecase(mockRepo)

	// Dados de entrada e entidade esperada
	input := dto.CreateOrderInput{
		Name:      "Test Order",
		TypeOrder: "Standard",
	}
	expectedOrder, _ := entity.NewOrder(input.Name, input.TypeOrder)

	// Definindo o comportamento do mock
	mockRepo.On("Save", expectedOrder).Return(nil)

	// Executando o teste
	err := usecaseCreate.Execute(&input)

	// Verificando o comportamento e os resultados
	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestCreateOrderUsecase_Execute_ErrorOnSave(t *testing.T) {
	// Configuração do mock e do caso de uso
	mockRepo := new(MockOrderRepository)
	usecaseCreate := usecase.NewCreateOrderUsecase(mockRepo)

	// Dados de entrada e entidade esperada
	input := dto.CreateOrderInput{
		Name:      "Test Order",
		TypeOrder: "Standard",
	}
	expectedOrder, _ := entity.NewOrder(input.Name, input.TypeOrder)

	// Definindo o comportamento do mock para erro ao salvar
	mockRepo.On("Save", expectedOrder).Return(errors.New("error saving order"))

	// Executando o teste
	err := usecaseCreate.Execute(&input)

	// Verificando o comportamento e os resultados
	assert.Error(t, err)
	assert.Equal(t, "error saving order", err.Error())
	mockRepo.AssertExpectations(t)
}

func TestCreateOrderUsecase_Execute_InvalidOrder(t *testing.T) {
	// Configuração do mock e do caso de uso
	mockRepo := new(MockOrderRepository)
	usecaseCreate := usecase.NewCreateOrderUsecase(mockRepo)

	// Dados de entrada inválidos
	input := dto.CreateOrderInput{
		Name:      "", // Nome vazio para simular erro de validação
		TypeOrder: "Standard",
	}

	// Executando o teste
	err := usecaseCreate.Execute(&input)

	// Verificando que o erro de validação é retornado
	assert.Error(t, err)
	assert.Equal(t, "order name is required", err.Error())
	mockRepo.AssertNotCalled(t, "Save", mock.Anything) // Certifica-se de que Save não foi chamado
}
