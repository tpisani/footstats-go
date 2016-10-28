package store

import (
	"sync"
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

func (s *Store) ParticipatingTeams() map[int]*footstats.Team {
	return s.teams
}

func (s *Store) TodaysMatches() map[int]*footstats.Match {
	return s.todaysMatches
}

func (s *Store) ClearGoals() {
	s.goalsIDs = []int{}
}

func (s *Store) LoadMatches() error {
	championships, err := s.client.Championships()
	if err != nil {
		return err
	}

	teams := make(map[int]*footstats.Team)
	todaysMatches := make(map[int]*footstats.Match)

	curYear, curMonth, curDay := time.Now().Date()

	wg := &sync.WaitGroup{}

	for _, championship := range championships {
		wg.Add(1)

		go func(c *footstats.Championship) {
			defer wg.Done()

			innerWg := &sync.WaitGroup{}
			innerWg.Add(2)

			var teamIDs []int

			go func() {

				defer innerWg.Done()

				matches, err := s.client.Matches(c.FootstatsID)
				if err != nil {
					return
				}

				for _, m := range matches {
					year, month, day := m.ScheduledTo.Date()
					if curYear == year && curMonth == month && curDay == day {
						todaysMatches[m.FootstatsID] = m
						teamIDs = append(teamIDs, m.HomeTeamID, m.VisitingTeamID)
					}
				}

			}()

			var entities *footstats.Entities

			go func() {

				defer innerWg.Done()

				e, err := s.client.Entities(c.FootstatsID)
				if err == nil {
					entities = e
				}

			}()

			innerWg.Wait()

			for _, t := range entities.Teams {
				for _, id := range teamIDs {
					if t.FootstatsID == id {
						teams[t.FootstatsID] = t
					}
				}
			}

		}(championship)

	}

	wg.Wait()

	s.teams = teams
	s.todaysMatches = todaysMatches

	return nil
}
