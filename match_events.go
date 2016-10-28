package footstats

import (
	"bytes"
	"encoding/json"
	"strconv"
)

type MatchEvents struct {
	HomeTeamScore     int
	VisitingTeamScore int

	Goals []*Goal
	Cards []*Card
}

type matchEvents struct {
	Championship struct {
		Match struct {
			HomeTeamScore     string `json:"PlacarMandante"`
			VisitingTeamScore string `json:"PlacarVisitante"`
			Goals             *struct {
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

func (m *MatchEvents) UnmarshalJSON(data []byte) error {
	var o matchEvents

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

	homeTeamScore, _ := strconv.Atoi(o.Championship.Match.HomeTeamScore)
	visitingTeamScore, _ := strconv.Atoi(o.Championship.Match.VisitingTeamScore)

	m.HomeTeamScore = homeTeamScore
	m.VisitingTeamScore = visitingTeamScore

	m.Goals = goals
	m.Cards = cards

	return nil
}
