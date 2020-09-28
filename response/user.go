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
	TeamName  string    `json:"teamname"`
	CreatedAt time.Time `json:"createdAt"`
	UpdateAt  time.Time `json:"updatedAt"`
}

type SingleTeamResponse struct {
	Team TeamResponse `json:"team"`
}

type TeamListResponse struct {
	Teams []*TeamResponse `json:"teams"`
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
