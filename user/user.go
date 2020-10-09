package user

import "github.com/HideBa/soroha-api/model"

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error

	CreateTeam(*model.Team, uint) error
	TeamByName(string) (*model.Team, error)
	TeamsList(uint, int) ([]model.Team, error)
	TeamUsersList(teamName string) (model.Team, []model.User, error)
	AddUserOnTeam(*model.Team, []string) ([]model.User, error)
}
