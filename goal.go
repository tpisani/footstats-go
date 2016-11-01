package footstats

import (
	"encoding/json"
	"strconv"
)

type Goal struct {
	ID         int         `json:"id"`
	PlayerID   int         `json:"player_id"`
	PlayerName string      `json:"player_name"`
	TeamID     int         `json:"team_id"`
	Period     MatchPeriod `json:"period"`
	Minute     int         `json:"minute"`
	Own        bool        `json:"own"`
}

type goal struct {
	ID     string  `json:"@Id"`
	Period string  `json:"@Periodo"`
	Minute string  `json:"@Momento"`
	Type   string  `json:"@Tipo"`
	Player *Player `json:"Jogador"`
	Team   *Team   `json:"Equipe"`
}

func (g *Goal) UnmarshalJSON(data []byte) error {
	var o goal

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)
	minute, _ := strconv.Atoi(o.Minute)

	g.ID = id
	g.Minute = minute
	g.PlayerID = o.Player.ID
	g.PlayerName = o.Player.Name
	g.TeamID = o.Team.ID

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
