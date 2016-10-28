package footstats

import (
	"encoding/json"
	"strconv"
)

type Goal struct {
	FootstatsID int
	PlayerID    int
	PlayerName  string
	TeamID      int
	Period      MatchPeriod
	Minute      int
	Own         bool
}

type goal struct {
	FootstatsID string  `json:"@Id"`
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

	footstatsID, _ := strconv.Atoi(o.FootstatsID)
	minute, _ := strconv.Atoi(o.Minute)

	g.FootstatsID = footstatsID
	g.Minute = minute
	g.PlayerID = o.Player.FootstatsID
	g.PlayerName = o.Player.Name
	g.TeamID = o.Team.FootstatsID

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
