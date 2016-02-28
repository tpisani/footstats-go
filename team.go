package footstats

import (
	"strconv"
)

type Team struct {
	FootstatsId   int64
	Name          string
	Initials      string
	IsPlaceholder bool
}

type footstatsTeam struct {
	FootstatsId   string `json:"@Id"`
	Name          string `json:"@Nome"`
	Initials      string `json:"@Sigla"`
	IsPlaceholder string `json:"@EquipeFantasia"`
}

func (f *footstatsTeam) team() *Team {
	footstatsId, _ := strconv.ParseInt(f.FootstatsId, 10, 64)
	isPlaceholder, _ := strconv.ParseBool(f.IsPlaceholder)

	return &Team{
		FootstatsId:   footstatsId,
		Name:          f.Name,
		Initials:      f.Initials,
		IsPlaceholder: isPlaceholder,
	}
}
