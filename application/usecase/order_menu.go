package usecase

import (
	"database/sql"
	"errors"

	"github.com/manomekun/ddd-sample00/consts"
	"github.com/manomekun/ddd-sample00/domain/model"
	"github.com/manomekun/ddd-sample00/infrastructure/infraservice/repository"
	"github.com/manomekun/ddd-sample00/prerror"
)

type OrderMenuParam struct {
	MenuID   uint
	Quantity uint
	UserID   string
}

func OrderMenu(param OrderMenuParam) error {
	ru := repository.NewUserRepository(false)
	user, err := ru.FindByID(param.UserID)
	if err != nil {
		return err
	}

	rm := repository.NewMenuRepository(false)
	if _, err := rm.FindByID(param.MenuID); err != nil {
		if err == sql.ErrNoRows {
			return errors.New(prerror.ErrNotFound)
		}
		return err
	}

	rom := repository.NewOrderMenuRepository(true)
	// model を先に作って渡すことで、model をそっくりそのまま作るようにrepositoryに指示する
	if err := rom.Create(model.Order{
		OrderType: consts.OrderTypeMenu,
		TargetID:  param.MenuID,
		Quantity:  param.Quantity,
		User:      *user,
	}); err != nil {
		return err
	}

	return nil
}

type ListOrderMenuParam struct {
	UserID string
}

func ListOrderMenu(param ListOrderMenuParam) ([]model.Order, error) {
	rm := repository.NewOrderMenuRepository(false)
	orders, err := rm.ListByUserID(param.UserID)
	if err != nil {
		return nil, err
	}

	return orders, nil
}
