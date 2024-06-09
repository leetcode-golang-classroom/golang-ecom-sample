package order

import (
	"database/sql"

	"github.com/leetcode-golang-classroom/golang-ecom-sample/types"
)

type Store struct {
	db *sql.DB
}

func NewStore(db *sql.DB) *Store {
	return &Store{db: db}
}

func (s *Store) CreateOrder(tx *sql.Tx, order types.Order) (int, error) {
	res, err := tx.Exec("INSERT INTO orders (userId, total, status, address) VAlUES (?,?,?,?)",
		order.UserID, order.Total, order.Status, order.Address,
	)
	if err != nil {
		return 0, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

func (s *Store) CreateOrderItem(tx *sql.Tx, orderItem types.OrderItem) error {
	_, err := tx.Exec("INSERT INTO order_items (orderId, productId, quantity, price) VALUES (?, ?, ?, ?)", orderItem.OrderID, orderItem.ProductID, orderItem.Quantity, orderItem.Price)
	return err
}
