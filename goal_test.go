package footstats

import (
	"testing"
)

var fgoal *footstatsGoal = &footstatsGoal{
	FootstatsId: "9999",
	Minute:      "23",
	Player: &footstatsPlayer{
		FootstatsId: "8888",
	},
	Team: &footstatsTeam{
		FootstatsId: "7777",
	},
}

func TestGoalMatchPeriod(t *testing.T) {
	var goal *Goal

	fgoal.Period = "Primeiro tempo"
	goal = fgoal.goal(9999)
	if goal.Period != FirstHalf {
		t.Errorf("Expected match period FirstHalf (%i) for 'Primeiro tempo', got %i",
			FirstHalf, goal.Period)
	}

	fgoal.Period = "Segundo tempo"
	goal = fgoal.goal(9999)
	if goal.Period != SecondHalf {
		t.Errorf("Expected match period SecondHalf (%i) for 'Segundo tempo', got %i",
			SecondHalf, goal.Period)
	}
}

func TestOwnGoal(t *testing.T) {
	var goal *Goal

	fgoal.Type = "Contra"
	goal = fgoal.goal(9999)
	if goal.Own != true {
		t.Error("Expected own as true, got", goal.Own)
	}

	fgoal.Type = "Favor"
	goal = fgoal.goal(9999)
	if goal.Own != false {
		t.Error("Expected own as false, got", goal.Own)
	}
}
