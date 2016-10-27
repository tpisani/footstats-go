package footstats

import (
	"encoding/json"
	"strconv"
)

type Player struct {
	FootstatsId int
	Name        string
	TeamId      int
}

type player struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
	TeamId      string `json:"@IdEquipe"`
}

func (p *Player) UnmarshalJSON(data []byte) error {
	var o player

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsId, _ := strconv.Atoi(o.FootstatsId)
	teamId, _ := strconv.Atoi(o.TeamId)

	p.FootstatsId = footstatsId
	p.Name = o.Name
	p.TeamId = teamId

	return nil
}
