package footstats

import (
	"encoding/json"
	"strconv"
)

type Player struct {
	FootstatsID int
	Name        string
	TeamID      int
}

type player struct {
	FootstatsID string `json:"@Id"`
	Name        string `json:"@Nome"`
	TeamID      string `json:"@IdEquipe"`
}

func (p *Player) UnmarshalJSON(data []byte) error {
	var o player

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsID, _ := strconv.Atoi(o.FootstatsID)
	teamID, _ := strconv.Atoi(o.TeamID)

	p.FootstatsID = footstatsID
	p.Name = o.Name
	p.TeamID = teamID

	return nil
}
