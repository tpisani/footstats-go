package footstats

import (
	"encoding/json"
	"strconv"
)

type Stadium struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	City          string `json:"city"`
	State         string `json:"state"`
	IsPlaceholder bool   `json:"is_placeholder"`
}

type stadium struct {
	ID    string `json:"@Id"`
	Name  string `json:"@Nome"`
	City  string `json:"@Cidade"`
	State string `json:"@Estado"`
}

func (s *Stadium) UnmarshalJSON(data []byte) error {
	var o stadium

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)

	s.ID = id
	s.Name = o.Name
	s.City = o.City
	s.State = o.State

	switch o.Name {
	case "A Definir":
		s.IsPlaceholder = true
	default:
		s.IsPlaceholder = false
	}

	return nil
}
