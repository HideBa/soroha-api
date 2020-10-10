package request

import (
	"github.com/HideBa/soroha-api/model"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type ExpenseCreateRequest struct {
	Expense struct {
		Price int `json:"price" validate:"required"`
		// UsedDate time.Time `json:"usedDate" validate:"required"`
		Comment string `json:"comment, omitempty"`
	} `json:"expense"`
}

type ExpenseUpdateRequest struct {
	Expense struct {
		Price   int    `json:"price"`
		Comment string `json:"comment"`
	}
}

func (req *ExpenseCreateRequest) Bind(c echo.Context, e *model.Expense) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	// if err := c.Validate(req); err != nil {
	// 	return err
	// }
	uuid, _ := uuid.NewUUID()
	// uuidStr := uuid
	e.Slug = uuid
	e.Price = req.Expense.Price
	// e.UsedDate = req.Expense.UsedDate
	e.Comment = req.Expense.Comment
	e.IsCalculated = false
	return nil
}

func (req *ExpenseUpdateRequest) ConvertModelToRequest(e *model.Expense) {
	req.Expense.Price = e.Price
	req.Expense.Comment = e.Comment
}

func (req *ExpenseUpdateRequest) Bind(c echo.Context, e *model.Expense) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	e.Price = req.Expense.Price
	e.Comment = req.Expense.Comment
	return nil
}

type CalculateExpensesRequest struct {
	TeamName string `json:"teamName"`
}

func (req *CalculateExpensesRequest) Bind(c echo.Context, calc *model.Calculation) error {
	if err := c.Bind(req); err != nil {
		return err
	}
	if err := c.Validate(req); err != nil {
		return err
	}
	return nil
}
