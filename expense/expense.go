package expense

import "github.com/HideBa/soroha-api/model"

type Store interface {
	CreateExpense(*model.Expense) error
	List(limit int) ([]model.Expense, int, error)
}
