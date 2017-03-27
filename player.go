package footstats

import (
	"encoding/json"
	"strconv"
	"strings"
)

type Player struct {
	ID       int
	TeamID   int
	Name     string
	Nickname string
	Position string
	Height   float32
}

type footstatsPlayer struct {
	ID       string `json:"Id"`
	TeamID   string `json:"IdEquipe"`
	Name     string `json:"Nome"`
	Nickname string `json:"Apelido"`
	Position string `json:"Posicao"`
	Height   string `json:"Altura"`
}

func (p *Player) UnmarshalJSON(data []byte) error {
	var o footstatsPlayer

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)
	teamID, _ := strconv.Atoi(o.TeamID)

	height, _ := strconv.ParseFloat(strings.Replace(o.Height, ",", ".", 1), 32)

	p.ID = id
	p.TeamID = teamID
	p.Name = o.Name
	p.Nickname = o.Nickname
	p.Position = o.Position
	p.Height = float32(height)

	return nil
}
