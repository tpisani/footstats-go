package footstats

import (
	"strconv"
	"time"
)

type MatchStatus int

const (
	NotStarted MatchStatus = iota
	OnGoing
	Finished
)

// https://golang.org/src/time/format.go
const footstatsDateLayout = "1/2/2006 3:04:05 PM"

// TODO: Team data, e.g. team IDs and score.
type Match struct {
	FootstatsId     int64
	Date            time.Time
	Status          MatchStatus
	Round           int
	StadiumId       int64
	RefereeId       int64
	HasLiveCoverage bool
}

type footstatsMatch struct {
	FootstatsId     string `json:"@Id"`
	Date            string `json:"Data"`
	Status          string
	Round           string `json:"Rodada"`
	StadiumId       string `json:"IdEstadio"`
	RefereeId       string `json:"IdArbitro"`
	HasLiveCoverage string `json:"AoVivo"`
}

type matchData struct {
	Data innerMatchData `json:"Partidas"`
}

func (m *matchData) innerData() []*footstatsMatch {
	return m.Data.Data
}

type innerMatchData struct {
	Data []*footstatsMatch `json:"Partida"`
}

func (m *matchData) matches() []*Match {
	var matches []*Match

	for _, d := range m.innerData() {
		footstatsId, _ := strconv.ParseInt(d.FootstatsId, 10, 64)
		date, _ := time.Parse(footstatsDateLayout, d.Date)
		round, _ := strconv.Atoi(d.Round)
		stadiumId, _ := strconv.ParseInt(d.StadiumId, 10, 64)
		refereeId, _ := strconv.ParseInt(d.RefereeId, 10, 64)

		var status MatchStatus
		switch d.Status {
		case "Partida em andamento":
			status = OnGoing
		case "Partida encerrada":
			status = Finished
		default:
			status = NotStarted
		}

		var hasLiveCoverage bool
		switch d.HasLiveCoverage {
		case "Sim":
			hasLiveCoverage = true
		default:
			hasLiveCoverage = false
		}

		matches = append(matches, &Match{
			FootstatsId:     footstatsId,
			Date:            date,
			Status:          status,
			Round:           round,
			StadiumId:       stadiumId,
			RefereeId:       refereeId,
			HasLiveCoverage: hasLiveCoverage,
		})
	}

	return matches
}
