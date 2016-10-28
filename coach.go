package footstats

import (
	"encoding/json"
	"strconv"
)

type Coach struct {
	FootstatsID int
	Name        string
}

type coach struct {
	FootstatsID string `json:"@Id"`
	Name        string `json:"@Nome"`
}

func (c *Coach) UnmarshalJSON(data []byte) error {
	var o coach

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsID, _ := strconv.Atoi(o.FootstatsID)

	c.FootstatsID = footstatsID
	c.Name = o.Name

	return nil
}
