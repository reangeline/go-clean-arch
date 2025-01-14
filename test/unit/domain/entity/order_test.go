package entity_test

import (
	"testing"

	"github.com/reangeline/go-clean-arch/internal/domain/entity"
	"github.com/stretchr/testify/assert"
)

func TestNewOrder(t *testing.T) {
	// Valid order creation
	order, err := entity.NewOrder("Order1", "Type1")
	assert.NoError(t, err)
	assert.NotNil(t, order)
	assert.Equal(t, "Order1", order.Name)
	assert.Equal(t, "Type1", order.TypeOrder)

	// Invalid order creation with missing name
	order, err = entity.NewOrder("", "Type1")
	assert.Error(t, err)
	assert.Equal(t, "order name is required", err.Error())
	assert.Nil(t, order)

	// Invalid order creation with missing type order
	order, err = entity.NewOrder("Order2", "")
	assert.Error(t, err)
	assert.Equal(t, "type order is required", err.Error())
	assert.Nil(t, order)
}

func TestIsValid(t *testing.T) {
	order := &entity.Order{
		Name:      "Order4",
		TypeOrder: "Type3",
	}

	// Test with valid data
	err := order.IsValid()
	assert.NoError(t, err)

	// Test with missing name
	order.Name = ""
	err = order.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "order name is required", err.Error())

	// Test with missing type order
	order.Name = "Order4"
	order.TypeOrder = ""
	err = order.IsValid()
	assert.Error(t, err)
	assert.Equal(t, "type order is required", err.Error())
}
