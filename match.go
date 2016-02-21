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

type Match struct {
	FootstatsId              int64
	Date                     time.Time
	Status                   MatchStatus
	Round                    int
	ChampionshipId           int64
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

type footstatsMatch struct {
	FootstatsId     string       `json:"@Id"`
	Date            string       `json:"Data"`
	Status          string       `json:"Status"`
	Round           string       `json:"Rodada"`
	Teams           []*matchTeam `json:"Equipe"`
	StadiumId       string       `json:"IdEstadio"`
	RefereeId       string       `json:"IdArbitro"`
	HasLiveCoverage string       `json:"AoVivo"`
}

type matchTeam struct {
	FootstatsId  string `json:"@Id"`
	Score        string `json:"@Placar"`
	PenaltyScore string `json:"@PlacarPenaltis"`
	Type         string `json:"@Tipo"`
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

func (m *matchData) matches(championshipId int64) []*Match {
	var matches []*Match

	for _, d := range m.innerData() {
		footstatsId, _ := strconv.ParseInt(d.FootstatsId, 10, 64)
		date, _ := time.Parse(footstatsDateLayout, d.Date)
		round, _ := strconv.Atoi(d.Round)
		stadiumId, _ := strconv.ParseInt(d.StadiumId, 10, 64)
		refereeId, _ := strconv.ParseInt(d.RefereeId, 10, 64)

		var status MatchStatus
		switch d.Status {
		case "Partida não iniciada":
			status = NotStarted
		case "Partida encerrada":
			status = Finished
		default:
			status = OnGoing
		}

		var hasLiveCoverage bool
		switch d.HasLiveCoverage {
		case "Sim":
			hasLiveCoverage = true
		default:
			hasLiveCoverage = false
		}

		var homeTeam, visitingTeam *matchTeam
		for _, t := range d.Teams {
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

		matches = append(matches, &Match{
			FootstatsId:              footstatsId,
			Date:                     date,
			Status:                   status,
			Round:                    round,
			ChampionshipId:           championshipId,
			HomeTeamId:               homeTeamId,
			HomeTeamScore:            homeTeamScore,
			HomeTeamPenaltyScore:     homeTeamPenaltyScore,
			VisitingTeamId:           visitingTeamId,
			VisitingTeamScore:        visitingTeamScore,
			VisitingTeamPenaltyScore: visitingTeamPenaltyScore,
			StadiumId:                stadiumId,
			RefereeId:                refereeId,
			HasLiveCoverage:          hasLiveCoverage,
		})
	}

	return matches
}
