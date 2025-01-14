package entity

import (
	"errors"

	pkg_entity "github.com/reangeline/go-clean-arch/pkg/entity"
)

type Order struct {
	OrderId   pkg_entity.ID `json:"order_id"`
	Name      string        `json:"name"`
	TypeOrder string        `json:"type_order"`
}

func NewOrder(name, type_order string) (*Order, error) {
	order := &Order{
		Name:      name,
		TypeOrder: type_order,
	}

	err := order.IsValid()
	if err != nil {
		return nil, err
	}

	return order, nil

}

func (o *Order) AddId() {
	o.OrderId = pkg_entity.NewID()
}

func (o *Order) IsValid() error {
	if o.Name == "" {
		return errors.New("order name is required")
	}

	if o.TypeOrder == "" {
		return errors.New("type order is required")
	}

	return nil
}
