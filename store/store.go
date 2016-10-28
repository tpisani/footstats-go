package store

import (
	"time"

	"footstats-go"
)

type Store struct {
	client *footstats.Client

	goalsIDs      []int
	teams         map[int]*footstats.Team
	todaysMatches map[int]*footstats.Match
}

func New(c *footstats.Client) *Store {
	return &Store{
		client: c,
	}
}

func (s *Store) TodaysMatches() map[int]*footstats.Match {
	return s.todaysMatches
}

func (s *Store) Hydrate() error {
	championships, err := s.client.Championships()
	if err != nil {
		return err
	}

	teams := make(map[int]*footstats.Team)
	todaysMatches := make(map[int]*footstats.Match)

	curYear, curMonth, curDay := time.Now().Date()

	for _, championship := range championships {
		matches, err := s.client.Matches(championship.FootstatsID)
		if err != nil {
			return err
		}

		var teamIDs []int

		for _, m := range matches {
			year, month, day := m.ScheduledTo.Date()
			if curYear == year && curMonth == month && curDay == day {
				todaysMatches[m.FootstatsID] = m
				teamIDs = append(teamIDs, m.HomeTeamID, m.VisitingTeamID)
			}
		}

		entities, err := s.client.Entities(championship.FootstatsID)
		if err != nil {
			return err
		}

		for _, t := range entities.Teams {
			for _, id := range teamIDs {
				if t.FootstatsID == id {
					teams[t.FootstatsID] = t
				}
			}
		}
	}

	s.goalsIDs = []int{}
	s.teams = teams
	s.todaysMatches = todaysMatches

	return nil
}
