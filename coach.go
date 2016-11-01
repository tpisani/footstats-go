package footstats

import (
	"encoding/json"
	"strconv"
)

type Coach struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type coach struct {
	ID   string `json:"@Id"`
	Name string `json:"@Nome"`
}

func (c *Coach) UnmarshalJSON(data []byte) error {
	var o coach

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)

	c.ID = id
	c.Name = o.Name

	return nil
}
