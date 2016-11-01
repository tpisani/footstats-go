package footstats

import (
	"encoding/json"
	"strconv"
)

type Championship struct {
	ID                int    `json:"id"`
	Name              string `json:"name"`
	HasClassification bool   `json:"has_classification"`
	CurrentRound      int    `json:"current_round"`
	TotalRounds       int    `json:"total_round"`
}

type championship struct {
	ID                string `json:"@Id"`
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

	id, _ := strconv.Atoi(o.ID)
	hasClassification, _ := strconv.ParseBool(o.HasClassification)
	currentRound, _ := strconv.Atoi(o.CurrentRound)
	totalRounds, _ := strconv.Atoi(o.TotalRounds)

	c.ID = id
	c.HasClassification = hasClassification
	c.CurrentRound = currentRound
	c.TotalRounds = totalRounds

	return nil
}
