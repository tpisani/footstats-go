package footstats

type Live struct {
	goals []*Goal
}

func (l *Live) Goals() []*Goal {
	return l.goals
}

type liveData struct {
	Campeonato struct {
		Partida struct {
			Gols *struct {
				Gol []*footstatsGoal
			}
		}
	}
}

func (l *liveData) goals(matchId int64) []*Goal {
	var goals []*Goal

	if l.Campeonato.Partida.Gols != nil {
		for _, d := range l.Campeonato.Partida.Gols.Gol {
			goals = append(goals, d.goal(matchId))
		}
	}

	return goals
}
