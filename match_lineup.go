package footstats

import (
	"encoding/json"
)

type MatchLineup struct {
	HomeTeamLineup     []Player
	VisitingTeamLineup []Player
}

type footstatsMatchLineup struct {
	HomeTeamLineup     []Player `json:"Mandante"`
	VisitingTeamLineup []Player `json:"Visitante"`
}

func (m *MatchLineup) UnmarshalJSON(data []byte) error {
	var o footstatsMatchLineup

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	m.HomeTeamLineup = o.HomeTeamLineup
	m.VisitingTeamLineup = o.VisitingTeamLineup

	return nil
}
