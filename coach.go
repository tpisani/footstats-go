package footstats

import (
	"encoding/json"
	"strconv"
)

type Coach struct {
	FootstatsId int
	Name        string
}

type coach struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
}

func (c *Coach) UnmarshalJSON(data []byte) error {
	var o coach

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsId, _ := strconv.Atoi(o.FootstatsId)

	c.FootstatsId = footstatsId
	c.Name = o.Name

	return nil
}
