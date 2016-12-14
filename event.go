package footstats

type GoalEvent struct {
	Match *Match
	Goal  *Goal
}

type CardEvent struct {
	Match *Match
	Card  *Card
}
