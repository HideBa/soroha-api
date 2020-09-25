package store

import (
	"github.com/HideBa/soroha-api/model"
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

func (expenseStore *ExpenseStore) CreateExpense(e *model.Expense) (err error) {
	//must implement transaction manually when expense has many to many relations with other models
	tx := expenseStore.db.Begin()
	if err := tx.Create(&e).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(e.ID).Preload("User").Find(&e).Error; err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
	// return expenseStore.db.Create(e).Error
}

func (expenseStore *ExpenseStore) GetList(limit int) ([]model.Expense, int, error) {
	var (
		expenses []model.Expense
		count    int
	)
	expenseStore.db.Model(&expenses).Count(&count)
	expenseStore.db.Preload("User").Limit(limit).Order("created_at desc").Find(&expenses)
	return expenses, count, nil
}
