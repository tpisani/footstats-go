package footstats

import (
	"encoding/json"
	"strconv"
)

type Stadium struct {
	FootstatsID   int
	Name          string
	City          string
	State         string
	IsPlaceholder bool
}

type stadium struct {
	FootstatsID string `json:"@Id"`
	Name        string `json:"@Nome"`
	City        string `json:"@Cidade"`
	State       string `json:"@Estado"`
}

func (s *Stadium) UnmarshalJSON(data []byte) error {
	var o stadium

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsID, _ := strconv.Atoi(o.FootstatsID)

	s.FootstatsID = footstatsID
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
