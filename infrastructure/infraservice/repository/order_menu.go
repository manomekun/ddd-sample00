package repository

import (
	"errors"
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/manomekun/ddd-sample00/consts"
	"github.com/manomekun/ddd-sample00/domain/model"
	"github.com/manomekun/ddd-sample00/infrastructure"
)

type OrderMenuRepository struct {
	DB *sqlx.DB
}

type OrderMenuWithUserDBModel struct {
	ID        uint      `db:"id"`
	MenuID    uint      `db:"menu_id"`
	MenuName  string    `db:"menu_name"`
	Quantity  uint      `db:"quantity"`
	Subtotal  uint      `db:"subtotal"`
	UserID    string    `db:"user_id"`
	UserName  string    `db:"user_name"`
	BirthDate time.Time `db:"birth_date"`
}

func NewOrderMenuRepository(writer bool) *OrderMenuRepository {
	var r OrderMenuRepository
	r.DB = infrastructure.DBReader
	if writer {
		r.DB = infrastructure.DB
	}
	return &r
}

func (r *OrderMenuRepository) Create(order model.Order) error {
	if order.OrderType != consts.OrderTypeMenu {
		return errors.New("invalid order type")
	}

	query := `INSERT INTO ...`
	if _, err := r.DB.NamedExec(query, map[string]interface{}{
		"menuID":   order.OrderType,
		"quantity": order.Quantity,
		"userID":   order.User.ID,
	}); err != nil {
		return err
	}
	return nil
}

func (r *OrderMenuRepository) ListByUserID(userID string) ([]model.Order, error) {
	var orders []*OrderMenuWithUserDBModel

	query := `
		SELECT om.id,
			om.menu_id,
			om.quantity,
			m.price,
			m.name AS menu_name,
			u.id AS user_id,
			u.name AS user_name,
			u.birth_date
		FROM orders_memus om
		JOIN menus m ON m.id = om.menu_id
		JOIN users u ON u.id = om.user_id
		JOIN payments p ON p.order_id = om.order_id
		WHERE om.ser_id = ?;`

	if err := r.DB.Select(&orders, query, userID); err != nil {
		return nil, err
	}

	// こういうロジックは repository/dto として分けてもいいかもね
	var res []model.Order
	for _, o := range orders {
		res = append(res, model.Order{
			ID:         o.ID,
			OrderType:  consts.OrderTypeMenu,
			TargetID:   o.MenuID,
			TargetName: o.MenuName,
			SubTotal:   o.Subtotal,
			Quantity:   o.Quantity,
			User: model.User{
				ID:        o.UserID,
				Name:      o.UserName,
				BirthDate: o.BirthDate,
			},
		})
	}

	return res, nil
}
