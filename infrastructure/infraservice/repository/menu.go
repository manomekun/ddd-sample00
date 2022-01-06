package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/manomekun/ddd-sample00/domain/model"
	"github.com/manomekun/ddd-sample00/infrastructure"
)

type MenuRepository struct {
	DB *sqlx.DB
}

type MenuDBModel struct {
	ID    uint   `db:"id"`
	Name  string `db:"name"`
	Price uint   `db:"price"`
}

func NewMenuRepository(writer bool) *MenuRepository {
	var r MenuRepository
	r.DB = infrastructure.DBReader
	if writer {
		r.DB = infrastructure.DB
	}
	return &r
}

func (r *MenuRepository) FindByID(id uint) (*model.Menu, error) {
	var m MenuDBModel

	query := `SELECT id, name, price FROM memus WHERE id = ?;`
	if err := r.DB.Get(&m, query, id); err != nil {
		return nil, err
	}

	var res model.Menu
	res.ID = m.ID
	res.Name = m.Name
	res.Price = m.Price

	return &res, nil
}
