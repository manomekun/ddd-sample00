package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/manomekun/ddd-sample00/domain/model"
	"github.com/manomekun/ddd-sample00/infrastructure"
)

type UserRepository struct {
	DB *sqlx.DB
}

type UserDBModel struct {
	ID        string    `db:"id"`
	Name      string    `db:"name"`
	BirthDate time.Time `db:"birth_date"`
}

func NewUserRepository(writer bool) *UserRepository {
	var r UserRepository
	r.DB = infrastructure.DBReader
	if writer {
		r.DB = infrastructure.DB
	}
	return &r
}

func (r *UserRepository) FindByID(id string) (*model.User, error) {
	var u UserDBModel

	query := `SELECT id, name, birth_date FROM users WHERE id = ?;`
	if err := r.DB.Get(&u, query, id); err != nil {
		return nil, err
	}

	var res model.User
	res.ID = u.ID
	res.Name = u.Name
	res.BirthDate = u.BirthDate

	return &res, nil
}
