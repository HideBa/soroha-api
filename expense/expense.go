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
	CalCulateExpenses([]model.Calculation, *model.Team, []model.User) ([]model.Calculation, error)
	UpdateCalculation(*model.Calculation) error
	GetCalculationBySlug(slug uuid.UUID) (*model.Calculation, error)
	CalculationsList(teamName string, calculations []model.Calculation) error
	ListByUser(uint, int, string) ([]model.Expense, int, error)
	// GetCalculation(*model.User) (*model.Calculation, error)
	ListByTeam(int, string) ([]model.Expense, int, error)
}
