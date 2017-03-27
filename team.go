package footstats

import (
	"encoding/json"
	"strconv"
)

type Team struct {
	ID            int
	Name          string
	Initials      string
	LogoURL       string
	IsPlaceholder bool
}

type footstatsTeam struct {
	ID            string `json:"Id"`
	Name          string `json:"Nome"`
	Initials      string `json:"Sigla"`
	LogoURL       string `json:"URLLogo"`
	IsPlaceholder string `json:"EquipeFantasia"`
}

func (t *Team) UnmarshalJSON(data []byte) error {
	var o footstatsTeam

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)
	isPlaceholder, _ := strconv.ParseBool(o.IsPlaceholder)

	t.ID = id
	t.Name = o.Name
	t.Initials = o.Initials
	t.LogoURL = o.LogoURL
	t.IsPlaceholder = isPlaceholder

	return nil
}
