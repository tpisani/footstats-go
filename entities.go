package footstats

type Entities struct {
	Teams    []*Team
	Players  []*Player
	Coaches  []*Coach
	Referees []*Referee
	Stadiums []*Stadium
}

type entitiesWrapper struct {
	Teams struct {
		Team []*Team `json:"Equipe"`
	} `json:"Equipes"`
	Players struct {
		Player []*Player `json:"Jogador"`
	} `json:"Jogadores"`
	Coaches struct {
		Coach []*Coach `json:"Tecnico"`
	} `json:"Tecnicos"`
	Referees struct {
		Referee []*Referee `json:"Arbitro"`
	} `json:"Arbitros"`
	Stadiums struct {
		Stadium []*Stadium `json:"Estadio"`
	} `json:"Estadios"`
}
