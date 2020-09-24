package response

import (
	"time"

	"github.com/HideBa/soroha-api/model"
	"github.com/labstack/echo/v4"
)

type expenseResponse struct {
	Price     int       `json:"price"`
	UsedDate  time.Time `json:"usedDate"`
	Comment   string    `json""comment"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt`
	User      struct {
		Username string `json:"username"`
	} `json:"user"`
}

type singleExpenseResponse struct {
	Expense expenseResponse `json:"expense"`
}

func NewExpenseResponse(c echo.Context, e *model.Expense) *singleExpenseResponse {
	expenseRes := expenseResponse{}
	expenseRes.Price = e.Price
	expenseRes.UsedDate = e.UsedDate
	expenseRes.Comment = e.Comment
	expenseRes.CreatedAt = e.CreatedAt
	expenseRes.UpdatedAt = e.UpdatedAt
	expenseRes.User.Username = e.User.Username
	return &singleExpenseResponse{expenseRes}
}
