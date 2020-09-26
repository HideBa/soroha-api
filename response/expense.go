package response

import (
	"time"

	"github.com/HideBa/soroha-api/model"
	"github.com/HideBa/soroha-api/user"
	"github.com/google/uuid"
	"github.com/labstack/echo/v4"
)

type expenseResponse struct {
	Slug  uuid.UUID `json:"slug"`
	Price int       `json:"price"`
	// UsedDate  time.Time `json:"usedDate"`
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

type expenseListResponse struct {
	Expenses      []*expenseResponse `json:"expenses"`
	ExpensesCount int                `json:"expensesCount"`
}

func NewExpenseResponse(c echo.Context, e *model.Expense) *singleExpenseResponse {
	expenseRes := expenseResponse{}
	expenseRes.Slug = e.Slug
	expenseRes.Price = e.Price
	// expenseRes.UsedDate = e.UsedDate
	expenseRes.Comment = e.Comment
	expenseRes.CreatedAt = e.CreatedAt
	expenseRes.UpdatedAt = e.UpdatedAt
	expenseRes.User.Username = e.User.Username
	return &singleExpenseResponse{expenseRes}
}

func NewExponseListResponse(us user.Store, userID uint, expenses []model.Expense, count int) *expenseListResponse {
	res := new(expenseListResponse)
	res.Expenses = make([]*expenseResponse, 0)
	for _, expense := range expenses {
		er := new(expenseResponse)
		er.Slug = expense.Slug
		er.Price = expense.Price
		er.Comment = expense.Comment
		er.CreatedAt = expense.CreatedAt
		er.UpdatedAt = expense.UpdatedAt
		er.User.Username = expense.User.Username
		res.Expenses = append(res.Expenses, er)
	}
	res.ExpensesCount = count
	return res
}
