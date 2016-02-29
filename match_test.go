package footstats

import (
	"testing"
)

func TestMatchStatus(t *testing.T) {
	fmatch := &footstatsMatch{
		FootstatsId: "9999",
		Teams: []*matchTeam{
			&matchTeam{
				FootstatsId:  "8888",
				Score:        "0",
				PenaltyScore: "0",
				Type:         "Mandante",
			},
			&matchTeam{
				FootstatsId:  "7777",
				Score:        "0",
				PenaltyScore: "0",
				Type:         "Visitante",
			},
		},
	}

	var match *Match

	fmatch.Status = "Partida não iniciada"
	match = fmatch.match(434)
	if match.Status != NotStarted {
		t.Errorf("Expected match status NotStarted (%i) for 'Partida não iniciada', got %i",
			NotStarted, match.Status)
	}

	fmatch.Status = "Partida encerrada"
	match = fmatch.match(434)
	if match.Status != Finished {
		t.Errorf("Expected match status Finished (%i) for 'Partida encerrada', got %i",
			Finished, match.Status)
	}

	fmatch.Status = "Primeiro tempo"
	match = fmatch.match(434)
	if match.Status != OnGoing {
		t.Errorf("Expected match status OnGoing (%i), got %i",
			OnGoing, match.Status)
	}

	fmatch.Status = "Partida cancelada"
	match = fmatch.match(434)
	if match.Status != Cancelled {
		t.Errorf("Expected match status Cancelled (%i) for 'Partida cancelada', got %i",
			Cancelled, match.Status)
	}
}
