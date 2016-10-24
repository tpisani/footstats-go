package footstats

import (
	"encoding/json"
	"strconv"
)

type Stadium struct {
	FootstatsId   int64
	Name          string
	City          string
	State         string
	IsPlaceholder bool
}

type stadium struct {
	FootstatsId string `json:"@Id"`
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

	footstatsId, _ := strconv.ParseInt(o.FootstatsId, 10, 64)

	s.FootstatsId = footstatsId
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
