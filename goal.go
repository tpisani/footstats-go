package footstats

import (
	"strconv"
)

type Goal struct {
	FootstatsId int64
	MatchId     int64
	PlayerId    int64
	TeamId      int64
	Period      MatchPeriod
	Minute      int
	Own         bool
}

type footstatsGoal struct {
	FootstatsId string           `json:"@Id"`
	Period      string           `json:"@Periodo"`
	Minute      string           `json:"@Momento"`
	Type        string           `json:"@Tipo"`
	Player      *footstatsPlayer `json:"Jogador"`
	Team        *footstatsTeam   `json:"Equipe"`
}

func (f *footstatsGoal) goal(matchId int64) *Goal {
	footstatsId, _ := strconv.ParseInt(f.FootstatsId, 10, 64)
	minute, _ := strconv.Atoi(f.FootstatsId)
	playerId, _ := strconv.ParseInt(f.Player.FootstatsId, 10, 64)
	teamId, _ := strconv.ParseInt(f.Team.FootstatsId, 10, 64)

	var period MatchPeriod
	switch f.Period {
	case "Primeiro tempo":
		period = FirstHalf
	case "Segundo tempo":
		period = SecondHalf
	}

	var own bool
	switch f.Type {
	case "Contra":
		own = true
	default:
		own = false
	}

	return &Goal{
		FootstatsId: footstatsId,
		MatchId:     matchId,
		PlayerId:    playerId,
		TeamId:      teamId,
		Period:      period,
		Minute:      minute,
		Own:         own,
	}
}
