package footstats

import (
	"strconv"
)

type Stadium struct {
	FootstatsId   int64
	Name          string
	City          string
	State         string
	IsPlaceholder bool
}

type footstatsStadium struct {
	FootstatsId string `json:"@Id"`
	Name        string `json:"@Nome"`
	City        string `json:"@Cidade"`
	State       string `json:"@Estado"`
}

func (f *footstatsStadium) stadium() *Stadium {
	footstatsId, _ := strconv.ParseInt(f.FootstatsId, 10, 64)

	var isPlaceholder bool
	switch f.Name {
	case "A Definir":
		isPlaceholder = true
	default:
		isPlaceholder = false
	}

	return &Stadium{
		FootstatsId:   footstatsId,
		Name:          f.Name,
		City:          f.City,
		State:         f.State,
		IsPlaceholder: isPlaceholder,
	}
}
