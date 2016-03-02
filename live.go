package footstats

import (
	"errors"
	"reflect"
	"strconv"
)

type Live struct {
	cards []*Card
	goals []*Goal
}

func (l *Live) Cards() []*Card {
	return l.cards
}

func (l *Live) Goals() []*Goal {
	return l.goals
}

type liveData struct {
	Campeonato struct {
		Partida struct {
			Gols *struct {
				Gol interface{}
			}
		}
		Stats struct {
			Equipe []struct {
				Cartoes *struct {
					Cartao interface{}
				}
			}
		}
	}
}

func cardFromData(data interface{}, matchId int64) (*Card, error) {
	dataMap, _ := data.(map[string]interface{})

	_, ok := dataMap["@IdCartao"]
	if !ok {
		return nil, errors.New("missing @IdCartao")
	}

	id := dataMap["@IdCartao"].(string)

	footstatsId, _ := strconv.ParseInt(id, 10, 64)
	playerId, _ := strconv.ParseInt(dataMap["@Id"].(string), 10, 64)
	minute, _ := strconv.Atoi(dataMap["@Minuto"].(string))

	var period MatchPeriod
	switch dataMap["@Periodo"].(string) {
	case "Primeiro tempo":
		period = FirstHalf
	case "Segundo tempo":
		period = SecondHalf
	}

	var cardType CardType
	switch dataMap["@Tipo"].(string) {
	case "Vermelho":
		cardType = RedCard
	case "Amarelo":
		cardType = YellowCard
	}

	return &Card{
		FootstatsId: footstatsId,
		MatchId:     matchId,
		PlayerId:    playerId,
		Period:      period,
		Minute:      minute,
		Type:        cardType,
	}, nil
}

func (l *liveData) cards(matchId int64) []*Card {
	var cards []*Card

	for _, e := range l.Campeonato.Stats.Equipe {
		if e.Cartoes == nil {
			continue
		}

		if reflect.TypeOf(e.Cartoes.Cartao).Kind() == reflect.Slice {
			s := reflect.ValueOf(e.Cartoes.Cartao)

			for i := 0; i < s.Len(); i++ {
				c, err := cardFromData(s.Index(i).Interface(), matchId)
				if err == nil {
					cards = append(cards, c)
				}

			}
		} else {
			c, err := cardFromData(e.Cartoes.Cartao, matchId)
			if err == nil {
				cards = append(cards, c)
			}
		}
	}

	return cards
}

func goalFromData(data interface{}, matchId int64) (*Goal, error) {
	dataMap := data.(map[string]interface{})

	footstatsId, _ := strconv.ParseInt(dataMap["@Id"].(string), 10, 64)
	minute, _ := strconv.Atoi(dataMap["@Momento"].(string))

	playerData := dataMap["Jogador"].(map[string]interface{})
	playerId, _ := strconv.ParseInt(playerData["@Id"].(string), 10, 64)

	var period MatchPeriod
	switch dataMap["@Periodo"].(string) {
	case "Primeiro tempo":
		period = FirstHalf
	case "Segundo tempo":
		period = SecondHalf
	}

	var own bool
	switch dataMap["@Tipo"].(string) {
	case "Contra":
		own = true
	default:
		own = false
	}

	return &Goal{
		FootstatsId: footstatsId,
		MatchId:     matchId,
		PlayerId:    playerId,
		Period:      period,
		Minute:      minute,
		Own:         own,
	}, nil
}

func (l *liveData) goals(matchId int64) []*Goal {
	var goals []*Goal

	goalsData := l.Campeonato.Partida.Gols
	if goalsData != nil {
		if reflect.TypeOf(goalsData.Gol).Kind() == reflect.Slice {
			s := reflect.ValueOf(goalsData.Gol)

			for i := 0; i < s.Len(); i++ {
				g, _ := goalFromData(s.Index(i).Interface(), matchId)
				goals = append(goals, g)
			}
		}
	}

	return goals
}
