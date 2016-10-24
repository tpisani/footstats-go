package footstats

import (
	"encoding/json"
	"strconv"
	"time"
)

// https://golang.org/src/time/format.go
const footstatsDateLayout = "1/2/2006 3:04:05 PM"

type MatchStatus int

const (
	NotStarted MatchStatus = iota
	InProgress
	Finished
	Cancelled
)

type MatchPeriod int

const (
	FirstHalf MatchPeriod = iota
	SecondHalf
)

type Match struct {
	FootstatsId              int64
	Date                     time.Time
	Status                   MatchStatus
	Round                    int
	HomeTeamId               int64
	HomeTeamScore            int
	HomeTeamPenaltyScore     int
	VisitingTeamId           int64
	VisitingTeamScore        int
	VisitingTeamPenaltyScore int
	StadiumId                int64
	RefereeId                int64
	HasLiveCoverage          bool
}

type matchWrapper struct {
	Matches struct {
		Match []*Match `json:"Partida"`
	} `json:"Partidas"`
}

type matchTeam struct {
	FootstatsId  string `json:"@Id"`
	Score        string `json:"@Placar"`
	PenaltyScore string `json:"@PlacarPenaltis"`
	Type         string `json:"@Tipo"`
}

type match struct {
	FootstatsId     string       `json:"@Id"`
	Date            string       `json:"Data"`
	Status          string       `json:"Status"`
	Round           string       `json:"Rodada"`
	Teams           []*matchTeam `json:"Equipe"`
	StadiumId       string       `json:"IdEstadio"`
	RefereeId       string       `json:"IdArbitro"`
	HasLiveCoverage string       `json:"AoVivo"`
}

func (m *Match) UnmarshalJSON(data []byte) error {
	var o match

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsId, _ := strconv.ParseInt(o.FootstatsId, 10, 64)
	date, _ := time.Parse(footstatsDateLayout, o.Date)
	round, _ := strconv.Atoi(o.Round)
	stadiumId, _ := strconv.ParseInt(o.StadiumId, 10, 64)
	refereeId, _ := strconv.ParseInt(o.RefereeId, 10, 64)

	var status MatchStatus
	switch o.Status {
	case "Partida não iniciada":
		status = NotStarted
	case "Partida encerrada":
		status = Finished
	case "Partida cancelada":
		status = Cancelled
	default:
		status = InProgress
	}

	var hasLiveCoverage bool
	switch o.HasLiveCoverage {
	case "Sim":
		hasLiveCoverage = true
	default:
		hasLiveCoverage = false
	}

	var homeTeam, visitingTeam *matchTeam
	for _, t := range o.Teams {
		if t.Type == "Mandante" {
			homeTeam = t
		} else {
			visitingTeam = t
		}
	}

	homeTeamId, _ := strconv.ParseInt(homeTeam.FootstatsId, 10, 64)
	homeTeamScore, _ := strconv.Atoi(homeTeam.Score)
	homeTeamPenaltyScore, _ := strconv.Atoi(homeTeam.PenaltyScore)

	visitingTeamId, _ := strconv.ParseInt(visitingTeam.FootstatsId, 10, 64)
	visitingTeamScore, _ := strconv.Atoi(visitingTeam.Score)
	visitingTeamPenaltyScore, _ := strconv.Atoi(visitingTeam.PenaltyScore)

	m.FootstatsId = footstatsId
	m.Date = date

	m.Status = status

	m.Round = round
	m.StadiumId = stadiumId
	m.RefereeId = refereeId

	m.HomeTeamId = homeTeamId
	m.HomeTeamScore = homeTeamScore
	m.HomeTeamPenaltyScore = homeTeamPenaltyScore

	m.VisitingTeamId = visitingTeamId
	m.VisitingTeamScore = visitingTeamScore
	m.VisitingTeamPenaltyScore = visitingTeamPenaltyScore

	m.HasLiveCoverage = hasLiveCoverage

	return nil
}
