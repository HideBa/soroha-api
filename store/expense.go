package store

import (
	"github.com/jinzhu/gorm"
)

type ExpenseStore struct {
	db *gorm.DB
}

func NewExpenseStore(db *gorm.DB) *ExpenseStore {
	return &ExpenseStore{
		db: db,
	}
}

func (expenseStore *ExpenseStore) Create(e *model.Expense) (err error) {
	return expenseStore.db.Create(e).Error
}

func (expenseStore *ExpenseStore) GetAll() (*model.Expense[], error) {
	
}