package store

import (
	"footstats-go"
)

type GoalEvent struct {
	MatchId int `json:"match_id"`

	HomeTeamScore     int `json:"home_team_score"`
	VisitingTeamScore int `json:"visiting_team_score"`

	Team *footstats.Team `json:"team"`

	Goal *footstats.Goal `json:"goal"`
}
