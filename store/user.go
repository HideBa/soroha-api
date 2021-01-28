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

func (userStore *UserStore) List() ([]model.User, error) {
	var (
		users []model.User
		err   error
	)
	err = userStore.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
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

func (userStore *UserStore) Create(u *model.User, t *model.Team) (err error) {
	tx := userStore.db.Begin()
	if err := userStore.db.Create(u).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := userStore.CreateTeam(t, u.ID); err != nil {
		tx.Rollback()
		return err
	}
	return tx.Commit().Error
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

func (userStore *UserStore) TeamUsersList(teamName string) (model.Team, []model.User, error) {
	var (
		users []model.User
		team  model.Team
	)

	err := userStore.db.Where("team_name = ?", teamName).First(&team).Error
	if err != nil {
		return team, nil, err
	}

	userStore.db.Model(&team).Association("Users").Find(&users)

	return team, users, nil
}

func (userStore *UserStore) TeamByName(teamName string) (*model.Team, error) {
	var teamModel model.Team
	if err := userStore.db.Where(&model.Team{TeamName: teamName}).First(&teamModel).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			return nil, nil
		}
		return nil, err
	}
	return &teamModel, nil
}

func (userStore *UserStore) AddUserOnTeam(team *model.Team, users []string) ([]model.User, error) {
	tx := userStore.db.Begin()
	userModels := make([]model.User, 0)

	for _, u := range users {
		user := model.User{Username: u}

		err := tx.Where(&user).First(&user).Error
		if err != nil && !gorm.IsRecordNotFoundError(err) {
			tx.Rollback()
			return nil, err
		}
		userModels = append(userModels, user)
	}
	if err := userStore.db.Model(team).Association("Users").Append(userModels).Error; err != nil {
		tx.Rollback()
		return nil, err
	}
	return userModels, tx.Commit().Error
}
