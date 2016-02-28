package footstats

import (
	"strconv"
)

type Referee struct {
	FootstatsId int64
	Name        string
}

type footstatsReferee struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
}

func (f *footstatsReferee) referee() *Referee {
	footstatsId, _ := strconv.ParseInt(f.FootstatsId, 10, 64)

	return &Referee{
		FootstatsId: footstatsId,
		Name:        f.Name,
	}
}
