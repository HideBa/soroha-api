package store

import (
	"fmt"

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
	fmt.Println("-------", u)
	return userStore.db.Create(u).Error
}

func (userStore *UserStore) Update(u *model.User) (err error) {
	return userStore.db.Model(u).Update(u).Error
}
