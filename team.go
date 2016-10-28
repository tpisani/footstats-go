package footstats

import (
	"encoding/json"
	"strconv"
)

type Team struct {
	FootstatsID   int
	Name          string
	Initials      string
	IsPlaceholder bool
}

type team struct {
	FootstatsID   string `json:"@Id"`
	Name          string `json:"@Nome"`
	Initials      string `json:"@Sigla"`
	IsPlaceholder string `json:"@EquipeFantasia"`
}

func (t *Team) UnmarshalJSON(data []byte) error {
	var o team

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsID, _ := strconv.Atoi(o.FootstatsID)
	isPlaceholder, _ := strconv.ParseBool(o.IsPlaceholder)

	t.FootstatsID = footstatsID
	t.Name = o.Name
	t.Initials = o.Initials
	t.IsPlaceholder = isPlaceholder

	return nil
}
