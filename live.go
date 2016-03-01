package footstats

import (
	"errors"
	"reflect"
	"strconv"
)

type Live struct {
	cards []*Card
}

func (l *Live) Cards() []*Card {
	return l.cards
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

type Card struct {
	FootstatsId int64
}

func cardFromData(data interface{}) (*Card, error) {
	dataMap, _ := data.(map[string]interface{})

	_, ok := dataMap["@IdCartao"]
	if !ok {
		return nil, errors.New("missing @IdCartao")
	}

	id := dataMap["@IdCartao"].(string)

	footstatsId, _ := strconv.ParseInt(id, 10, 64)

	return &Card{
		FootstatsId: footstatsId,
	}, nil
}

func (l *liveData) cards() []*Card {
	var cards []*Card

	for _, e := range l.Campeonato.Stats.Equipe {
		if e.Cartoes == nil {
			continue
		}

		if reflect.TypeOf(e.Cartoes.Cartao).Kind() == reflect.Slice {
			s := reflect.ValueOf(e.Cartoes.Cartao)

			for i := 0; i < s.Len(); i++ {
				c, err := cardFromData(s.Index(i).Interface())
				if err == nil {
					cards = append(cards, c)
				}

			}
		} else {
			c, err := cardFromData(e.Cartoes.Cartao)
			if err == nil {
				cards = append(cards, c)
			}
		}
	}

	return cards
}
