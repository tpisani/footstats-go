package footstats

import (
	"encoding/json"
	"strconv"
)

type Player struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	TeamID int    `json:"team_id"`
}

type player struct {
	ID     string `json:"@Id"`
	Name   string `json:"@Nome"`
	TeamID string `json:"@IdEquipe"`
}

func (p *Player) UnmarshalJSON(data []byte) error {
	var o player

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)
	teamID, _ := strconv.Atoi(o.TeamID)

	p.ID = id
	p.Name = o.Name
	p.TeamID = teamID

	return nil
}
