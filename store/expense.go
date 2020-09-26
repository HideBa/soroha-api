package store

import (
	"github.com/HideBa/soroha-api/model"
	"github.com/google/uuid"
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

func (expenseStore *ExpenseStore) List(limit int) ([]model.Expense, int, error) {
	var (
		expenses []model.Expense
		count    int
	)
	expenseStore.db.Model(&expenses).Count(&count)
	expenseStore.db.Preload("User").Limit(limit).Order("created_at desc").Find(&expenses)
	return expenses, count, nil
}

func (expenseStore *ExpenseStore) GetUserExpenseBySlug(userID uint, slug uuid.UUID) (*model.Expense, error) {
	var expenseModel model.Expense

	err := expenseStore.db.Where(&model.Expense{Slug: slug, UserID: userID}).Find(&expenseModel).Error
	if err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &expenseModel, err
}

func (expenseStore *ExpenseStore) UpdateExpense(e *model.Expense) (err error) {
	tx := expenseStore.db.Begin()
	if err := tx.Model(e).Update(e).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(e.ID).Preload("User").Find(e).Error; err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit().Error
}

func (expenseStore *ExpenseStore) DeleteExpense(e *model.Expense) (err error) {
	return expenseStore.db.Delete(e).Error
}

// func (expenseStore *ExpenseStore) ListByUser(userID uint, limit int)
