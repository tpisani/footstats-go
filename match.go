package footstats

import (
	"encoding/json"
	"strconv"
	"time"
)

// https://golang.org/src/time/format.go
const footstatsDateLayout = "02/01/2006 15:04:05"

type MatchPeriod int

const (
	NotStarted MatchPeriod = iota

	FirstHalf
	FirstInterval

	SecondHalf
	SecondInterval

	FirstExtraHalf
	ThirdInterval

	SecondExtraHalf

	Penalties

	Finished

	Interrupted
	Canceled
)

type Match struct {
	ID                       int
	ChampionshipID           int
	Round                    int
	HomeTeamID               int
	HomeTeamName             string
	HomeTeamScore            int
	HomeTeamPenaltyScore     int
	VisitingTeamID           int
	VisitingTeamName         string
	VisitingTeamScore        int
	VisitingTeamPenaltyScore int
	Period                   MatchPeriod
	Date                     time.Time
	Goals                    []Goal
}

type footstatsMatch struct {
	ID                       string `json:"Id"`
	ChampionshipID           string `json:"IdCampeonato"`
	Round                    string `json:"Rodada"`
	HomeTeamID               string `json:"IdEquipeMandante"`
	HomeTeamName             string `json:"NomeMandante"`
	HomeTeamScore            string `json:"PlacarMandante"`
	HomeTeamPenaltyScore     string `json:"PlacarPenaltisMandante"`
	VisitingTeamID           string `json:"IdEquipeVisitante"`
	VisitingTeamName         string `json:"NomeVisitante"`
	VisitingTeamScore        string `json:"PlacarVisitante"`
	VisitingTeamPenaltyScore string `json:"PlacarPenaltisVisitante"`
	Period                   string `json:"PeriodoAtual"`
	Date                     string `json:"Data"`
	Goals                    []Goal `json:"Gols"`
}

func matchPeriodFromString(periodString string) MatchPeriod {
	var period MatchPeriod

	switch periodString {
	case "Partida não iniciada":
		period = NotStarted
	case "Primeiro tempo":
		period = FirstHalf
	case "Intervalo 1":
		period = FirstInterval
	case "Segundo tempo":
		period = SecondHalf
	case "Intervalo 2":
		period = SecondInterval
	case "Prorrogação 1":
		period = FirstExtraHalf
	case "Intervalo 3":
		period = ThirdInterval
	case "Prorrogação 2":
		period = SecondExtraHalf
	case "Disputa de Pênaltis":
		period = Penalties
	case "Partida encerrada":
		period = Finished
	case "Partida interrompida":
		period = Interrupted
	case "Partida cancelada":
		period = Canceled
	}

	return period
}

func (m *Match) UnmarshalJSON(data []byte) error {
	var o footstatsMatch

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)
	championshipID, _ := strconv.Atoi(o.ChampionshipID)
	round, _ := strconv.Atoi(o.Round)

	homeTeamID, _ := strconv.Atoi(o.HomeTeamID)
	homeTeamScore, _ := strconv.Atoi(o.HomeTeamScore)
	homeTeamPenaltyScore, _ := strconv.Atoi(o.HomeTeamPenaltyScore)

	visitingTeamID, _ := strconv.Atoi(o.VisitingTeamID)
	visitingTeamScore, _ := strconv.Atoi(o.VisitingTeamScore)
	visitingTeamPenaltyScore, _ := strconv.Atoi(o.VisitingTeamPenaltyScore)

	period := matchPeriodFromString(o.Period)

	date, _ := time.Parse(footstatsDateLayout, o.Date)

	m.ID = id
	m.ChampionshipID = championshipID
	m.Round = round

	m.HomeTeamID = homeTeamID
	m.HomeTeamScore = homeTeamScore
	m.HomeTeamName = o.HomeTeamName
	m.HomeTeamPenaltyScore = homeTeamPenaltyScore

	m.VisitingTeamID = visitingTeamID
	m.VisitingTeamName = o.VisitingTeamName
	m.VisitingTeamScore = visitingTeamScore
	m.VisitingTeamPenaltyScore = visitingTeamPenaltyScore

	m.Period = period

	m.Date = date

	m.Goals = o.Goals

	return nil
}
