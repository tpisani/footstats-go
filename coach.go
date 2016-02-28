package footstats

import (
	"strconv"
)

type Coach struct {
	FootstatsId int64
	Name        string
}

type footstatsCoach struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
}

func (f *footstatsCoach) coach() *Coach {
	footstatsId, _ := strconv.ParseInt(f.FootstatsId, 10, 64)

	return &Coach{
		FootstatsId: footstatsId,
		Name:        f.Name,
	}
}
