package store

import (
	"footstats-go"
)

type GoalEvent struct {
	HomeTeamScore     int
	VisitingTeamScore int

	Match *footstats.Match
	Team  *footstats.Team
	Goal  *footstats.Goal
}
