package footstats

import (
	"encoding/json"
	"strconv"
)

type Referee struct {
	FootstatsID int
	Name        string
}

type referee struct {
	FootstatsID string `json:"@Id"`
	Name        string `json:"@Nome"`
}

func (r *Referee) UnmarshalJSON(data []byte) error {
	var o referee

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsID, _ := strconv.Atoi(o.FootstatsID)

	r.FootstatsID = footstatsID
	r.Name = o.Name

	return nil
}
