package footstats

import (
	"encoding/json"
	"strconv"
)

type Goal struct {
	FootstatsId int
	PlayerId    int
	TeamId      int
	Period      MatchPeriod
	Minute      int
	Own         bool
}

type goal struct {
	FootstatsId string  `json:"@Id"`
	Period      string  `json:"@Periodo"`
	Minute      string  `json:"@Momento"`
	Type        string  `json:"@Tipo"`
	Player      *Player `json:"Jogador"`
	Team        *Team   `json:"Equipe"`
}

func (g *Goal) UnmarshalJSON(data []byte) error {
	var o goal

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsId, _ := strconv.Atoi(o.FootstatsId)
	minute, _ := strconv.Atoi(o.Minute)

	g.FootstatsId = footstatsId
	g.Minute = minute
	g.PlayerId = o.Player.FootstatsId
	g.TeamId = o.Team.FootstatsId

	switch o.Period {
	case "Primeiro tempo":
		g.Period = FirstHalf
	case "Segundo tempo":
		g.Period = SecondHalf
	}

	switch o.Type {
	case "Contra":
		g.Own = true
	default:
		g.Own = false
	}

	return nil
}
