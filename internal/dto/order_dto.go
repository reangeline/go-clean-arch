package dto

type CreateOrderInput struct {
	Name      string `json:"name"`
	TypeOrder string `json:"type_order"`
}

type ListOrdersOutput struct {
	OrderId   string `json:"order_id"`
	Name      string `json:"name"`
	TypeOrder string `json:"type_order"`
}
