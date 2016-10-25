package footstats

import (
	"encoding/json"
	"strconv"
)

type Player struct {
	FootstatsId int64
	Name        string
	TeamId      int64
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

	footstatsId, _ := strconv.ParseInt(o.FootstatsId, 10, 64)
	teamId, _ := strconv.ParseInt(o.TeamId, 10, 64)

	p.FootstatsId = footstatsId
	p.Name = o.Name
	p.TeamId = teamId

	return nil
}
