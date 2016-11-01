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
	ID       int         `json:"id"`
	PlayerID int         `json:"player_id"`
	Minute   int         `json:"minute"`
	Period   MatchPeriod `json:"period"`
	Type     CardType    `json:"type"`
}

type card struct {
	ID       string `json:"@IdCartao"`
	PlayerID string `json:"@Id"`
	Minute   string `json:"@Minuto"`
	Period   string `json:"@Periodo"`
	Type     string `json:"@Tipo"`
}

func (c *Card) UnmarshalJSON(data []byte) error {
	var o card

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)
	playerID, _ := strconv.Atoi(o.PlayerID)
	minute, _ := strconv.Atoi(o.Minute)

	c.ID = id
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
