package footstats

import (
	"strconv"
)

type Championship struct {
	FootstatsId       int64
	Name              string
	HasClassification bool
	CurrentRound      int
	Rounds            int
}

type footstatsChampionship struct {
	FootstatsId       string `json:"@Id"`
	Name              string `json:"@Nome"`
	HasClassification string `json:"@TemClassificacao"`
	CurrentRound      string `json:"@RodadaAtual"`
	Rounds            string `json:"@Rodadas"`
}

func (f *footstatsChampionship) championship() *Championship {
	footstatsId, _ := strconv.ParseInt(f.FootstatsId, 10, 64)
	hasClassification, _ := strconv.ParseBool(f.HasClassification)
	currentRound, _ := strconv.Atoi(f.CurrentRound)
	rounds, _ := strconv.Atoi(f.Rounds)

	return &Championship{
		FootstatsId:       footstatsId,
		Name:              f.Name,
		HasClassification: hasClassification,
		CurrentRound:      currentRound,
		Rounds:            rounds,
	}
}

type championshipData struct {
	Campeonato []*footstatsChampionship
}

func (c *championshipData) championships() []*Championship {
	var championships []*Championship

	for _, d := range c.Campeonato {
		championships = append(championships, d.championship())
	}

	return championships
}
