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
	FootstatsID int
	PlayerID    int
	Minute      int
	Period      MatchPeriod
	Type        CardType
}

type card struct {
	FootstatsID string `json:"@IdCartao"`
	PlayerID    string `json:"@Id"`
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

	footstatsID, _ := strconv.Atoi(o.FootstatsID)
	playerID, _ := strconv.Atoi(o.PlayerID)
	minute, _ := strconv.Atoi(o.Minute)

	c.FootstatsID = footstatsID
	c.PlayerID = playerID
	c.Minute = minute

	switch o.Type {
	case "Vermelho":
		c.Type = RedCard
	default:
		c.Type = YellowCard
	}

	return nil
}
