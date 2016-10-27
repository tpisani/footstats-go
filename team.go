package footstats

import (
	"encoding/json"
	"strconv"
)

type Team struct {
	FootstatsId   int
	Name          string
	Initials      string
	IsPlaceholder bool
}

type team struct {
	FootstatsId   string `json:"@Id"`
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

	footstatsId, _ := strconv.Atoi(o.FootstatsId)
	isPlaceholder, _ := strconv.ParseBool(o.IsPlaceholder)

	t.FootstatsId = footstatsId
	t.Name = o.Name
	t.Initials = o.Initials
	t.IsPlaceholder = isPlaceholder

	return nil
}
