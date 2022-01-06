package handler

import (
	"net/http"

	echo "github.com/labstack/echo/v4"
	"github.com/manomekun/ddd-sample00/application/usecase"
	"github.com/manomekun/ddd-sample00/prerror"
)

type ListOrderMenuResponse struct {
	ID       uint   `json:"id"`
	Quantity uint   `json:"quantity"`
	Price    uint   `json:"price"`
	Name     string `json:"name"`
}

func ListOrderMenu(c echo.Context) error {
	userID := c.QueryParam("userId")

	orders, err := usecase.ListOrderMenu(usecase.ListOrderMenuParam{
		UserID: userID,
	})
	if err != nil {
		if err.Error() == prerror.ErrNotFound {
			return c.JSON(http.StatusNotFound, "not found")
		}
	}
	var res ListOrderMenuResponse
	for _, o := range orders {
		res.ID = o.ID
		res.Quantity = o.Quantity
		res.Price = o.SubTotal
		res.Name = o.TargetName
	}

	return c.JSON(http.StatusOK, res)
}
