package footstats

import (
	"encoding/json"
	"strconv"
)

type Goal struct {
	ID         int
	PlayerID   int
	PlayerName string
	TeamID     int
	TeamName   string
	Minute     int
	Period     MatchPeriod
	Own        bool
}

type footstatsGoal struct {
	PlayerName string `json:"Jogador"`
	TeamName   string `json:"Equipe"`
	Period     string `json:"Periodo"`
	Minute     string `json:"Momento"`
	Type       string `json:"Tipo"`
}

func (g *Goal) UnmarshalJSON(data []byte) error {
	var o footstatsGoal

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	minute, _ := strconv.Atoi(o.Minute)

	period := matchPeriodFromString(o.Period)

	var own bool
	switch o.Type {
	case "Favor":
		own = false
	case "Contra":
		own = true
	}

	g.PlayerName = o.PlayerName
	g.TeamName = o.TeamName
	g.Period = period
	g.Minute = minute
	g.Own = own

	return nil
}
