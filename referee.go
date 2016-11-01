package footstats

import (
	"encoding/json"
	"strconv"
)

type Referee struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type referee struct {
	ID   string `json:"@Id"`
	Name string `json:"@Nome"`
}

func (r *Referee) UnmarshalJSON(data []byte) error {
	var o referee

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)

	r.ID = id
	r.Name = o.Name

	return nil
}
