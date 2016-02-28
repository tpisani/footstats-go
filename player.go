package footstats

import (
	"strconv"
)

type Player struct {
	FootstatsId int64
	Name        string
	TeamId      int64
}

type footstatsPlayer struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
	TeamId      string `json:"@IdEquipe"`
}

func (f *footstatsPlayer) player() *Player {
	footstatsId, _ := strconv.ParseInt(f.FootstatsId, 10, 64)
	teamId, _ := strconv.ParseInt(f.TeamId, 10, 64)

	return &Player{
		FootstatsId: footstatsId,
		Name:        f.Name,
		TeamId:      teamId,
	}
}
