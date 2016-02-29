package footstats

import (
	"strconv"
)

type PlayerStats struct {
	PlayerId int64

	OnField            bool
	SubstitutePlayerId int64

	Position string

	SuccessfulPasses      int
	SuccessfulCrosses     int
	SuccessfulLaunches    int
	SuccessfulDribbles    int
	SuccessfulDisarms     int
	SuccessfulCompletions int

	FailedPasses      int
	FailedCrosses     int
	FailedLaunches    int
	FaliedDribbles    int
	FailedDisarms     int
	FailedCompletions int

	FaultsCommited int
	FaultsReceived int

	PenaltiesCommited int
	PenaltiesReceived int

	Offsides int
	Defenses int
	Services int
}

type footstatsPlayerStats struct {
	PlayerId string `json:"@Id"`

	Status             string `json:"@Status"`
	SubstitutePlayerId string `json:"@IdSubstituto"`

	Position string `json:"@Posicao"`

	SuccessfulPasses      string `json:"@Passe_Certo"`
	SuccessfulCrosses     string `json:"Cruzamento_Certo"`
	SuccessfulLaunches    string `json:"Lancamento_Certo"`
	SuccessfulDribbles    string `json:"Drible_Certo"`
	SuccessfulDisarms     string `json:"Desarme_Certo"`
	SuccessfulCompletions string `json:"Finalizacao_Certa"`

	FailedPasses      string `json:"@Passe_Errado"`
	FailedCrosses     string `json:"Cruzamento_Errado"`
	FailedLaunches    string `json:"Lancamento_Errado"`
	FaliedDribbles    string `json:"Drible_Errado"`
	FailedDisarms     string `json:"Desarme_Errado"`
	FailedCompletions string `json:"Finalizacao_Errada"`

	FaultsCommited string `json:"Falta_Cometida"`
	FaultsReceived string `json:"Falta_Recebida"`

	PenaltiesCommited string `json:"Penalti_Cometido"`
	PenaltiesReceived string `json:"Penaitl_Recebido"`

	Offsides string `json:"Impedimento"`
	Defenses string `json:"Defesa"`
	Services string `json:"Assitencia"`
}

func (f *footstatsPlayerStats) playerStats() *PlayerStats {
	playerId, _ := strconv.ParseInt(f.PlayerId, 10, 64)
	substitutePlayerId, err := strconv.ParseInt(f.PlayerId, 10, 64)
	if err != nil {
		substitutePlayerId = 0
	}

	successfulPasses, _ := strconv.Atoi(f.SuccessfulPasses)
	successfulCrosses, _ := strconv.Atoi(f.SuccessfulCrosses)
	successfulLaunches, _ := strconv.Atoi(f.SuccessfulLaunches)
	successfulDribbles, _ := strconv.Atoi(f.SuccessfulDribbles)
	successfulDisarms, _ := strconv.Atoi(f.SuccessfulDisarms)
	successfulCompletions, _ := strconv.Atoi(f.SuccessfulCompletions)

	failedPasses, _ := strconv.Atoi(f.FailedPasses)
	failedCrosses, _ := strconv.Atoi(f.FailedCrosses)
	failedLaunches, _ := strconv.Atoi(f.FailedLaunches)
	faliedDribbles, _ := strconv.Atoi(f.FaliedDribbles)
	failedDisarms, _ := strconv.Atoi(f.FailedDisarms)
	failedCompletions, _ := strconv.Atoi(f.FailedCompletions)

	faultsCommited, _ := strconv.Atoi(f.FaultsCommited)
	faultsReceived, _ := strconv.Atoi(f.FaultsReceived)

	penaltiesCommited, _ := strconv.Atoi(f.PenaltiesCommited)
	penaltiesReceived, _ := strconv.Atoi(f.PenaltiesReceived)

	offsides, _ := strconv.Atoi(f.Offsides)
	defenses, _ := strconv.Atoi(f.Defenses)
	services, _ := strconv.Atoi(f.Services)

	var onField bool
	switch f.Status {
	case "EmCampo":
		onField = true
	default:
		onField = false
	}

	return &PlayerStats{
		PlayerId: playerId,

		OnField:            onField,
		SubstitutePlayerId: substitutePlayerId,

		SuccessfulPasses:      successfulPasses,
		SuccessfulCrosses:     successfulCrosses,
		SuccessfulLaunches:    successfulLaunches,
		SuccessfulDribbles:    successfulDribbles,
		SuccessfulDisarms:     successfulDisarms,
		SuccessfulCompletions: successfulCompletions,

		FailedPasses:      failedPasses,
		FailedCrosses:     failedCrosses,
		FailedLaunches:    failedLaunches,
		FaliedDribbles:    faliedDribbles,
		FailedDisarms:     failedDisarms,
		FailedCompletions: failedCompletions,

		FaultsCommited: faultsCommited,
		FaultsReceived: faultsReceived,

		PenaltiesCommited: penaltiesCommited,
		PenaltiesReceived: penaltiesReceived,

		Offsides: offsides,
		Defenses: defenses,
		Services: services,
	}
}
