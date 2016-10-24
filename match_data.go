package footstats

import (
	"bytes"
	"encoding/json"
)

type MatchData struct {
	Goals []*Goal
	Cards []*Card
}

type matchData struct {
	Championship struct {
		Match struct {
			Goals *struct {
				Goal *json.RawMessage `json:"Gol"`
			} `json:"Gols"`
		} `json:"Partida"`
		Stats struct {
			Team []struct {
				Cards *struct {
					Card *json.RawMessage `json:"Cartao"`
				} `json:"Cartoes"`
			} `json:"Equipe"`
		}
	} `json:"Campeonato"`
}

func (m *MatchData) UnmarshalJSON(data []byte) error {
	var o matchData

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	var goals []*Goal

	goalData := o.Championship.Match.Goals
	if goalData != nil {
		d := json.NewDecoder(bytes.NewReader(*goalData.Goal))

		t, err := d.Token()
		if err != nil {
			return err
		}

		if t == json.Delim('[') {
			json.Unmarshal(*goalData.Goal, &goals)
		} else {
			var goal *Goal
			json.Unmarshal(*goalData.Goal, &goal)
			goals = append(goals, goal)
		}
	}

	var cards []*Card

	for _, team := range o.Championship.Stats.Team {
		if team.Cards == nil {
			continue
		}

		d := json.NewDecoder(bytes.NewReader(*team.Cards.Card))

		t, err := d.Token()
		if err != nil {
			return err
		}

		if t == json.Delim('[') {
			json.Unmarshal(*team.Cards.Card, &cards)
		} else {
			var card *Card
			json.Unmarshal(*team.Cards.Card, &card)
			cards = append(cards, card)
		}
	}

	m.Goals = goals
	m.Cards = cards

	return nil
}
