package footstats

import (
	"strconv"
)

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
		footstatsId, _ := strconv.ParseInt(d.FootstatsId, 10, 64)
		isPlaceholder, _ := strconv.ParseBool(d.IsPlaceholder)

		teams = append(teams, &Team{
			FootstatsId:   footstatsId,
			Name:          d.Name,
			Initials:      d.Initials,
			IsPlaceholder: isPlaceholder,
		})
	}

	return teams
}

func (e *entitiesData) coaches() []*Coach {
	var coaches []*Coach

	for _, d := range e.Tecnicos.Tecnico {
		footstatsId, _ := strconv.ParseInt(d.FootstatsId, 10, 64)

		coaches = append(coaches, &Coach{
			FootstatsId: footstatsId,
			Name:        d.Name,
		})
	}

	return coaches
}

func (e *entitiesData) referees() []*Referee {
	var referees []*Referee

	for _, d := range e.Arbitros.Arbitro {
		footstatsId, _ := strconv.ParseInt(d.FootstatsId, 10, 64)

		referees = append(referees, &Referee{
			FootstatsId: footstatsId,
			Name:        d.Name,
		})
	}

	return referees
}

func (e *entitiesData) stadiums() []*Stadium {
	var stadiums []*Stadium

	for _, d := range e.Estadios.Estadio {
		footstatsId, _ := strconv.ParseInt(d.FootstatsId, 10, 64)

		var isPlaceholder bool
		switch d.Name {
		case "A Definir":
			isPlaceholder = true
		default:
			isPlaceholder = false
		}

		stadiums = append(stadiums, &Stadium{
			FootstatsId:   footstatsId,
			Name:          d.Name,
			City:          d.City,
			State:         d.State,
			IsPlaceholder: isPlaceholder,
		})
	}

	return stadiums
}
