package footstats

import (
	"encoding/json"
	"strconv"
)

type Team struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Initials      string `json:"initials"`
	IsPlaceholder bool   `json:"is_placeholder"`
}

type team struct {
	ID            string `json:"@Id"`
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

	id, _ := strconv.Atoi(o.ID)
	isPlaceholder, _ := strconv.ParseBool(o.IsPlaceholder)

	t.ID = id
	t.Name = o.Name
	t.Initials = o.Initials
	t.IsPlaceholder = isPlaceholder

	return nil
}
