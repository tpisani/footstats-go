package footstats

import (
	"encoding/json"
	"strconv"
)

type Championship struct {
	ID           int
	Name         string
	Rounds       int
	CurrentRound int
}

type footstatsChampionship struct {
	ID           string `json:"Id"`
	Name         string `json:"Nome"`
	Rounds       string `json:"QtdRodadas"`
	CurrentRound string `json:"RodadaAtual"`
}

func (c *Championship) UnmarshalJSON(data []byte) error {
	var o footstatsChampionship

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)
	rounds, _ := strconv.Atoi(o.Rounds)
	currentRound, _ := strconv.Atoi(o.CurrentRound)

	c.ID = id
	c.Name = o.Name
	c.Rounds = rounds
	c.CurrentRound = currentRound

	return nil
}
