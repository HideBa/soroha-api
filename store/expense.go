package store

import (
	"time"

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
	var team model.Team
	err = expenseStore.db.Where(model.Team{TeamName: e.Team.TeamName}).First(&team).Error
	if err != nil {
		return err
	}
	e.Team = team
	tx := expenseStore.db.Begin()
	if err := tx.Create(&e).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Where(e.ID).Preload("User").Find(&e).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Where(e.ID).Preload("Team").Find(&e).Error; err != nil {
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

func (expenseStore *ExpenseStore) ListByUser(userID uint, limit int, teamName string) ([]model.Expense, int, error) {
	var (
		user     model.User
		team     model.Team
		expenses []model.Expense
		count    int
	)
	err := expenseStore.db.First(&user, userID).Error

	if err != nil {
		return nil, 0, err
	}
	err = expenseStore.db.Where(&model.Team{TeamName: teamName}).First(&team).Error
	if err != nil {
		return nil, 0, err
	}
	expenseStore.db.Where(&model.Expense{UserID: user.ID, TeamID: team.ID}).Preload("User").Preload("Team").Limit(limit).Order("created_at").Find(&expenses)
	// expenseStore.db.Where(&model.Expense{UserID: user.ID}).Model(&model.User{}).Count(&count)
	count = len(expenses)

	return expenses, count, nil
}

func (expenseStore *ExpenseStore) ListByTeam(limit int, teamName string) ([]model.Expense, int, error) {
	var (
		team     model.Team
		expenses []model.Expense
		count    int
	)
	err := expenseStore.db.Where(&model.Team{TeamName: teamName}).First(&team).Error
	if err != nil {
		return nil, 0, err
	}
	err = expenseStore.db.Where(&model.Expense{TeamID: team.ID}).Preload("Team").Limit(limit).Order("created_at").Find(&expenses).Error
	if err != nil {
		return nil, 0, err
	}
	count = len(expenses)
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

func (expenseStore *ExpenseStore) CalCulateExpenses(calcs []model.Calculation, team *model.Team, users []model.User) ([]model.Calculation, error) {
	totalExpense, err := expensesTotal(expenseStore, team, users)
	if err != nil {
		return nil, err
	}
	expensePerUser := totalExpense / len(users)
	for _, user := range users {
		calc := model.Calculation{}
		uuid, _ := uuid.NewUUID()
		calc.Slug = uuid
		calc.CalculatedAt = time.Now()
		calc.IsPaid = false
		calc.Team = *team
		calc.User = user
		calc.Price = expensePerUser
		calcs = append(calcs, calc)
	}
	for _, calc := range calcs {
		if err := expenseStore.db.Create(&calc).Error; err != nil {
			return calcs, err
		}
	}
	return calcs, nil
}

func expensesTotal(expenseStore *ExpenseStore, team *model.Team, users []model.User) (total int, err error) {
	var expenses []model.Expense
	err = expenseStore.db.Preload("Team").Where(model.Expense{IsCalculated: false, TeamID: team.ID}).Find(&expenses).Error
	if err != nil {
		return 0, err
	}
	var expenseSum int
	for _, expense := range expenses {
		expenseSum += expense.Price
	}
	return expenseSum, nil
}

func (expenseStore *ExpenseStore) GetCalculationBySlug(slug uuid.UUID) (*model.Calculation, error) {
	var calculation model.Calculation
	err := expenseStore.db.Where(model.Calculation{Slug: slug}).Find(&calculation).Error
	if err != nil {
		return nil, err
	}
	return &calculation, nil
}

func (expenseStore *ExpenseStore) UpdateCalculation(cl *model.Calculation) error {
	err := expenseStore.db.Model(cl).Update(cl).Error
	if err != nil {
		return err
	}
	return nil
}
