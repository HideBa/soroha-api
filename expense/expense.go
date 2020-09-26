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
}
