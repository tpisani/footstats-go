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

type championshipData struct {
	Data []*footstatsChampionship `json:"Campeonato"`
}

func (c *championshipData) championships() []*Championship {
	var championships []*Championship

	for _, d := range c.Data {
		footstatsId, _ := strconv.ParseInt(d.FootstatsId, 10, 64)
		hasClassification, _ := strconv.ParseBool(d.HasClassification)
		currentRound, _ := strconv.Atoi(d.CurrentRound)
		rounds, _ := strconv.Atoi(d.Rounds)

		championships = append(championships, &Championship{
			FootstatsId:       footstatsId,
			Name:              d.Name,
			HasClassification: hasClassification,
			CurrentRound:      currentRound,
			Rounds:            rounds,
		})
	}

	return championships
}
