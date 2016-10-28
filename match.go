package footstats

import (
	"encoding/json"
	"strconv"
	"time"
)

// https://golang.org/src/time/format.go
const footstatsTimeLayout = "1/2/2006 3:04:05 PM"

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
	FootstatsID              int
	ScheduledTo              time.Time
	Status                   MatchStatus
	Round                    int
	HomeTeamID               int
	HomeTeamScore            int
	HomeTeamPenaltyScore     int
	VisitingTeamID           int
	VisitingTeamScore        int
	VisitingTeamPenaltyScore int
	StadiumID                int
	RefereeID                int
	HasLiveCoverage          bool
}

type matchWrapper struct {
	Matches struct {
		Match []*Match `json:"Partida"`
	} `json:"Partidas"`
}

type matchTeam struct {
	FootstatsID  string `json:"@Id"`
	Score        string `json:"@Placar"`
	PenaltyScore string `json:"@PlacarPenaltis"`
	Type         string `json:"@Tipo"`
}

type match struct {
	FootstatsID     string       `json:"@Id"`
	ScheduledTo     string       `json:"Data"`
	Status          string       `json:"Status"`
	Round           string       `json:"Rodada"`
	Teams           []*matchTeam `json:"Equipe"`
	StadiumID       string       `json:"IdEstadio"`
	RefereeID       string       `json:"IdArbitro"`
	HasLiveCoverage string       `json:"AoVivo"`
}

func (m *Match) UnmarshalJSON(data []byte) error {
	var o match

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	footstatsID, _ := strconv.Atoi(o.FootstatsID)
	scheduledTo, _ := time.Parse(footstatsTimeLayout, o.ScheduledTo)
	round, _ := strconv.Atoi(o.Round)
	stadiumID, _ := strconv.Atoi(o.StadiumID)
	refereeID, _ := strconv.Atoi(o.RefereeID)

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

	homeTeamID, _ := strconv.Atoi(homeTeam.FootstatsID)
	homeTeamScore, _ := strconv.Atoi(homeTeam.Score)
	homeTeamPenaltyScore, _ := strconv.Atoi(homeTeam.PenaltyScore)

	visitingTeamID, _ := strconv.Atoi(visitingTeam.FootstatsID)
	visitingTeamScore, _ := strconv.Atoi(visitingTeam.Score)
	visitingTeamPenaltyScore, _ := strconv.Atoi(visitingTeam.PenaltyScore)

	m.FootstatsID = footstatsID
	m.ScheduledTo = scheduledTo

	m.Status = status

	m.Round = round
	m.StadiumID = stadiumID
	m.RefereeID = refereeID

	m.HomeTeamID = homeTeamID
	m.HomeTeamScore = homeTeamScore
	m.HomeTeamPenaltyScore = homeTeamPenaltyScore

	m.VisitingTeamID = visitingTeamID
	m.VisitingTeamScore = visitingTeamScore
	m.VisitingTeamPenaltyScore = visitingTeamPenaltyScore

	m.HasLiveCoverage = hasLiveCoverage

	return nil
}
