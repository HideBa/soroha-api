package store

import (
	"github.com/HideBa/soroha-api/model"
	"github.com/jinzhu/gorm"
)

type UserStore struct {
	db *gorm.DB
}

func NewUserStore(db *gorm.DB) *UserStore {
	return &UserStore{
		db: db,
	}
}

func (userStore *UserStore) GetByID(id uint) (*model.User, error) {
	var m model.User
	if err := userStore.db.First(&m, id).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (userStore *UserStore) GetByUsername(username string) (*model.User, error) {
	var m model.User
	if err := userStore.db.Where(&model.User{Username: username}).First(&m).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &m, nil
}

func (userStore *UserStore) Create(u *model.User) (err error) {
	return userStore.db.Create(u).Error
}

func (userStore *UserStore) Update(u *model.User) (err error) {
	return userStore.db.Model(u).Update(u).Error
}

func (userStore *UserStore) CreateTeam(teamModel *model.Team, userID uint) (err error) {
	user := model.User{}
	user.ID = userID
	if err := userStore.db.Create(teamModel).Error; err != nil {
		return err
	}
	return userStore.db.Model(teamModel).Association("Users").Append(&user).Error
}

func (userStore *UserStore) TeamsList(userID uint, limit int) ([]model.Team, error) {
	var (
		user  model.User
		teams []model.Team
	)
	err := userStore.db.First(&user, userID).Error
	if err != nil {
		return nil, err
	}

	userStore.db.Preload("Users", "id = ?", userID).Find(&teams)

	return teams, nil
}
