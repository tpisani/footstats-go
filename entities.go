package footstats

type Entities struct {
	teams    []*Team
	coaches  []*Coach
	referees []*Referee
	stadiums []*Stadium
}

func (e *Entities) Teams() []*Team {
	return e.teams
}

func (e *Entities) Coaches() []*Coach {
	return e.coaches
}

func (e *Entities) Referees() []*Referee {
	return e.referees
}

func (e *Entities) Stadiums() []*Stadium {
	return e.stadiums
}

type entitiesData struct {
	Equipes struct {
		Equipe []*footstatsTeam
	}
	Arbitros struct {
		Arbitro []*footstatsReferee
	}
	Tecnicos struct {
		Tecnico []*footstatsCoach
	}
	Estadios struct {
		Estadio []*footstatsStadium
	}
}

func (e *entitiesData) teams() []*Team {
	var teams []*Team

	for _, d := range e.Equipes.Equipe {
		teams = append(teams, d.team())
	}

	return teams
}

func (e *entitiesData) coaches() []*Coach {
	var coaches []*Coach

	for _, d := range e.Tecnicos.Tecnico {
		coaches = append(coaches, d.coach())
	}

	return coaches
}

func (e *entitiesData) referees() []*Referee {
	var referees []*Referee

	for _, d := range e.Arbitros.Arbitro {
		referees = append(referees, d.referee())
	}

	return referees
}

func (e *entitiesData) stadiums() []*Stadium {
	var stadiums []*Stadium

	for _, d := range e.Estadios.Estadio {
		stadiums = append(stadiums, d.stadium())
	}

	return stadiums
}
