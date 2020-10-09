package response

import (
	"time"

	"github.com/HideBa/soroha-api/model"
	"github.com/HideBa/soroha-api/user"
	util "github.com/HideBa/soroha-api/utils"
)

type UserResponse struct {
	User struct {
		Username string `json:"username"`
		Token    string `json:"token"`
	} `json:"user"`
}

type TeamResponse struct {
	TeamName  string    `json:"teamName"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt"`
}

type SingleTeamResponse struct {
	Team TeamResponse `json:"team"`
}

type TeamListResponse struct {
	Teams []*TeamResponse `json:"teams"`
}

type TeamUsersResponse struct {
	Team  TeamResponse    `json:"team"`
	Users []*UserResponse `json:"users"`
}

func NewUserResponse(userModel *model.User) *UserResponse {
	res := &UserResponse{}
	res.User.Username = userModel.Username
	res.User.Token = util.GenerateJWT(userModel.ID)
	return res
}

func NewTeamResponse(teamModel *model.Team) *SingleTeamResponse {
	res := &SingleTeamResponse{}
	res.Team.TeamName = teamModel.TeamName
	return res
}

func TeamsListResponse(userStore user.Store, userID uint, teams []model.Team) *TeamListResponse {
	res := new(TeamListResponse)
	res.Teams = make([]*TeamResponse, 0)
	for _, team := range teams {
		teamRes := new(TeamResponse)
		teamRes.TeamName = team.TeamName
		teamRes.CreatedAt = team.CreatedAt
		teamRes.UpdateAt = team.UpdatedAt
		res.Teams = append(res.Teams, teamRes)
	}
	return res
}

func TeamUsersListResponse(team *model.Team, users []model.User) *TeamUsersResponse {
	res := new(TeamUsersResponse)
	res.Team.TeamName = team.TeamName
	res.Team.CreatedAt = team.CreatedAt
	res.Team.UpdateAt = team.UpdatedAt
	res.Users = make([]*UserResponse, 0)
	for _, user := range users {
		userRes := new(UserResponse)
		userRes.User.Username = user.Username
		res.Users = append(res.Users, userRes)
	}
	return res
}
