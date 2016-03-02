package footstats

type CardType int

const (
	RedCard CardType = iota
	YellowCard
)

type Card struct {
	FootstatsId int64
	MatchId     int64
	PlayerId    int64
	Period      MatchPeriod
	Minute      int
	Type        CardType
}
