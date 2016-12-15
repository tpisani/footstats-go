package footstats

type GoalEvent struct {
	Match *Match
	Goal  *Goal
}

type CardEvent struct {
	Match *Match
	Card  *Card
}

type MatchStatusEvent struct {
	Match      *Match
	UpdateType MatchStatusUpdateType
}

type MatchStatusUpdateType int

const (
	MatchStarted = iota
	MatchFinished
)
