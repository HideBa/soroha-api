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
	Comment      string `json:"comment"`
	IsCalculated bool   `json:"isCalculated"`
	CreatedAt    int64  `json:"createdAt"`
	UpdatedAt    int64  `json:"updatedAt"`
	User         struct {
		Username string `json:"username"`
	} `json:"user"`
}

type singleExpenseResponse struct {
	Expense expenseResponse `json:"expense"`
	Team    TeamResponse    `json:"team"`
}

type expenseListResponse struct {
	Expenses      []*expenseResponse `json:"expenses"`
	ExpensesCount int                `json:"expensesCount"`
}

func NewExpenseResponse(c echo.Context, e *model.Expense) *singleExpenseResponse {
	expenseRes := singleExpenseResponse{}
	expenseRes.Expense.Slug = e.Slug
	expenseRes.Expense.Price = e.Price
	// expenseRes.UsedDate = e.UsedDate
	expenseRes.Expense.Comment = e.Comment
	expenseRes.Expense.IsCalculated = e.IsCalculated
	expenseRes.Expense.CreatedAt = e.CreatedAt.Unix()
	expenseRes.Expense.UpdatedAt = e.UpdatedAt.Unix()
	expenseRes.Expense.User.Username = e.User.Username
	expenseRes.Team.TeamName = e.Team.TeamName
	// expenseRes.Team.CreatedAt = e.Team.CreatedAt
	// expenseRes.Team.UpdateAt = e.Team.UpdatedAt
	return &expenseRes
}

func ExpenseListResponse(us user.Store, expenses []model.Expense, count int) *expenseListResponse {
	res := new(expenseListResponse)
	res.Expenses = make([]*expenseResponse, 0)
	for _, expense := range expenses {
		er := new(expenseResponse)
		er.Slug = expense.Slug
		er.Price = expense.Price
		er.Comment = expense.Comment
		er.IsCalculated = expense.IsCalculated
		er.CreatedAt = expense.CreatedAt.Unix()
		er.UpdatedAt = expense.UpdatedAt.Unix()
		er.User.Username = expense.User.Username
		res.Expenses = append(res.Expenses, er)
	}
	res.ExpensesCount = count
	return res
}

type CalculationReseponse struct {
	CaluculatedAt time.Time `json:"calculatedAt"`
	Slug          uuid.UUID `json:"slug"`
	Price         int       `json:"price"`
	IsPaid        bool      `json:"isPaid"`
	UsersName     string    `json:"usersName"`
	TeamName      string    `json:"teamName"`
}

type SingleCalculationResponse struct {
	Calculation CalculationReseponse `json:"calculation"`
}

type CalculationsListResponse struct {
	calculations []CalculationReseponse `json:"calculations"`
}

func NewSingleCalculationResponse(c echo.Context, calc *model.Calculation) *SingleCalculationResponse {
	calcRes := &SingleCalculationResponse{}
	calcRes.Calculation.Slug = calc.Slug
	calcRes.Calculation.CaluculatedAt = calc.CalculatedAt
	calcRes.Calculation.IsPaid = calc.IsPaid
	calcRes.Calculation.Price = calc.Price
	calcRes.Calculation.UsersName = calc.User.Username
	calcRes.Calculation.TeamName = calc.Team.TeamName
	return calcRes
}

func NewCalculationsListResponse(c echo.Context, calculations []model.Calculation) *CalculationsListResponse {
	calculationsRes := &CalculationsListResponse{}
	for _, calc := range calculations {
		calcRes := CalculationReseponse{}
		calcRes.Slug = calc.Slug
		calcRes.Price = calc.Price
		calcRes.IsPaid = calc.IsPaid
		calcRes.TeamName = calc.Team.TeamName
		calcRes.UsersName = calc.User.Username
		calcRes.CaluculatedAt = calc.CalculatedAt
		calculationsRes.calculations = append(calculationsRes.calculations, calcRes)
	}
	return calculationsRes
}
