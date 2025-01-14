package repository

import (
	"database/sql"
	"fmt"

	"github.com/reangeline/go-clean-arch/internal/domain/entity"
)

type OrderRepository struct {
	dbConn *sql.DB
}

func NewOrderRepository(dbConn *sql.DB) *OrderRepository {
	return &OrderRepository{
		dbConn,
	}
}

func (o *OrderRepository) Save(order *entity.Order) error {

	query := `INSERT INTO orders (order_id, name, type_order) VALUES ($1, $2, $3)`
	_, err := o.dbConn.Exec(query, order.OrderId.String(), order.Name, order.TypeOrder)
	if err != nil {
		return fmt.Errorf("failed to insert order: %w", err)
	}
	return nil
}

func (o *OrderRepository) ListOrders() ([]*entity.Order, error) {

	query := `SELECT order_id, name, type_order FROM orders`
	rows, err := o.dbConn.Query(query)
	if err != nil {
		return nil, fmt.Errorf("failed to list orders: %w", err)
	}
	defer rows.Close()

	var orders []*entity.Order
	for rows.Next() {
		var order entity.Order
		var idStore string
		err := rows.Scan(&idStore, &order.Name, &order.TypeOrder)
		if err != nil {
			return nil, fmt.Errorf("failed to scan order row: %w", err)
		}

		orders = append(orders, &order)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return orders, nil
}
