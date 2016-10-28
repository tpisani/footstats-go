package footstats

import (
	"encoding/json"
	"strconv"
)

type Championship struct {
	FootstatsID       int
	Name              string
	HasClassification bool
	CurrentRound      int
	TotalRounds       int
}

type championship struct {
	FootstatsID       string `json:"@Id"`
	Name              string `json:"@Nome"`
	HasClassification string `json:"@TemClassificacao"`
	CurrentRound      string `json:"@RodadaAtual"`
	TotalRounds       string `json:"@Rodadas"`
}

type championshipWrapper struct {
	Championships []*Championship `json:"Campeonato"`
}

func (c *Championship) UnmarshalJSON(data []byte) error {
	var o championship

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsID, _ := strconv.Atoi(o.FootstatsID)
	hasClassification, _ := strconv.ParseBool(o.HasClassification)
	currentRound, _ := strconv.Atoi(o.CurrentRound)
	totalRounds, _ := strconv.Atoi(o.TotalRounds)

	c.FootstatsID = footstatsID
	c.HasClassification = hasClassification
	c.CurrentRound = currentRound
	c.TotalRounds = totalRounds

	return nil
}
