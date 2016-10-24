package footstats

import (
	"encoding/json"
	"strconv"
)

type Referee struct {
	FootstatsId int64
	Name        string
}

type referee struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
}

func (r *Referee) UnmarshalJSON(data []byte) error {
	var o referee

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsId, _ := strconv.ParseInt(o.FootstatsId, 10, 64)

	r.FootstatsId = footstatsId
	r.Name = o.Name

	return nil
}
