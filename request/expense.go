package request

import (
	"time"

	"github.com/HideBa/soroha-api/model"
	"github.com/labstack/echo/v4"
)

type ExpenseCreateRequest struct {
	Expense struct {
		Price    int       `json:"price" validate:"required"`
		UsedDate time.Time `json:"usedDate" validate:"required"`
		Comment  string    `json:"comment, omitempty"`
	} `json:"expense"`
}

func (req *ExpenseCreateRequest) Bind(c echo.Context, e *model.Expense) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	// if err := c.Validate(req); err != nil {
	// 	return err
	// }
	e.Price = req.Expense.Price
	e.UsedDate = req.Expense.UsedDate
	return nil
}
