package footstats

import (
	"encoding/json"
	"strconv"
)

type Referee struct {
	FootstatsId int
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

	footstatsId, _ := strconv.Atoi(o.FootstatsId)

	r.FootstatsId = footstatsId
	r.Name = o.Name

	return nil
}
