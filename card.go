package footstats

import (
	"encoding/json"
	"strconv"
)

type CardType int

const (
	RedCard CardType = iota
	YellowCard
)

type Card struct {
	FootstatsId int
	PlayerId    int
	Minute      int
	Period      MatchPeriod
	Type        CardType
}

type card struct {
	FootstatsId string `json:"@IdCartao"`
	PlayerId    string `json:"@Id"`
	Minute      string `json:"@Minuto"`
	Period      string `json:"@Periodo"`
	Type        string `json:"@Tipo"`
}

func (c *Card) UnmarshalJSON(data []byte) error {
	var o card

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsId, _ := strconv.Atoi(o.FootstatsId)
	playerId, _ := strconv.Atoi(o.PlayerId)
	minute, _ := strconv.Atoi(o.Minute)

	c.FootstatsId = footstatsId
	c.PlayerId = playerId
	c.Minute = minute

	switch o.Type {
	case "Vermelho":
		c.Type = RedCard
	default:
		c.Type = YellowCard
	}

	return nil
}
