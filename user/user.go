package user

import "github.com/HideBa/soroha-api/model"

type Store interface {
	GetByID(uint) (*model.User, error)
	GetByUsername(string) (*model.User, error)
	Create(*model.User) error
	Update(*model.User) error

	CreateTeam(*model.Team, uint) error
	TeamsList(uint, int) ([]model.Team, error)
}
