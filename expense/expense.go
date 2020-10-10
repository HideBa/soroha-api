package expense

import (
	"github.com/HideBa/soroha-api/model"
	"github.com/google/uuid"
)

type Store interface {
	CreateExpense(*model.Expense) error
	List(limit int) ([]model.Expense, int, error)
	GetUserExpenseBySlug(userID uint, slug uuid.UUID) (*model.Expense, error)
	UpdateExpense(*model.Expense) error
	DeleteExpense(*model.Expense) error
	CalCulateExpenses(*model.Calculation, *model.Team, []model.User) error
	GetCalculation(*model.User) (*model.Calculation, error)
}
