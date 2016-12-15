package footstats

import (
	"sync"
)

type Subscriber struct {
	initialized bool
	stopped     bool

	client  *Client
	matches []*Match

	goalMutex  sync.Mutex
	goalIDs    []int
	goalEvents chan *GoalEvent

	cardMutex  sync.Mutex
	cardIDs    []int
	cardEvents chan *CardEvent

	matchStatusEvents chan *MatchStatusEvent
}

func NewSubscriber(c *Client, matches []*Match) *Subscriber {
	s := &Subscriber{
		client: c,

		matches: matches,

		goalEvents:        make(chan *GoalEvent),
		cardEvents:        make(chan *CardEvent),
		matchStatusEvents: make(chan *MatchStatusEvent),
	}

	return s
}

// TODO: Goals and cards registry sure need some refactoring.

func (s *Subscriber) addGoal(g *Goal) bool {
	s.goalMutex.Lock()
	defer s.goalMutex.Unlock()

	for _, id := range s.goalIDs {
		if g.ID == id {
			return false
		}
	}

	s.goalIDs = append(s.goalIDs, g.ID)
	return true
}

func (s *Subscriber) addCard(c *Card) bool {
	s.cardMutex.Lock()
	defer s.cardMutex.Unlock()

	for _, id := range s.cardIDs {
		if c.ID == id {
			return false
		}
	}

	s.cardIDs = append(s.cardIDs, c.ID)
	return true
}

func (s *Subscriber) checkGoalUpdates(m *Match, stats *MatchStats) {
	for _, g := range stats.Goals {
		if s.addGoal(g) && s.initialized {
			m.HomeTeamScore = stats.HomeTeamScore
			m.VisitingTeamScore = stats.VisitingTeamScore

			s.goalEvents <- &GoalEvent{
				Match: m,
				Goal:  g,
			}
		}
	}
}

func (s *Subscriber) checkCardUpdates(m *Match, stats *MatchStats) {
	for _, c := range stats.Cards {
		if s.addCard(c) && s.initialized {
			s.cardEvents <- &CardEvent{
				Match: m,
				Card:  c,
			}
		}
	}
}

func (s *Subscriber) checkMatchStatusUpdates(m *Match, stats *MatchStats) {
	wasInProgress := m.Status == InProgress
	isInProgress := stats.Status == InProgress

	m.Status = stats.Status

	if !wasInProgress && isInProgress {
		s.matchStatusEvents <- &MatchStatusEvent{
			Match:      m,
			UpdateType: MatchStarted,
		}
	} else if wasInProgress && !isInProgress {
		s.matchStatusEvents <- &MatchStatusEvent{
			Match:      m,
			UpdateType: MatchFinished,
		}
	}
}

func (s *Subscriber) poll(wg *sync.WaitGroup, m *Match) {
	defer wg.Done()

	stats, err := s.client.MatchStats(m.ID)
	if err != nil {
		return
	}

	s.checkGoalUpdates(m, stats)
	s.checkCardUpdates(m, stats)
	s.checkMatchStatusUpdates(m, stats)
}

func (s *Subscriber) startPolling() {
	wg := &sync.WaitGroup{}

	for !s.stopped {
		for _, m := range s.matches {
			wg.Add(1)
			go s.poll(wg, m)
		}

		wg.Wait()
		s.initialized = true
	}
}

func (s *Subscriber) Start() {
	go s.startPolling()
}

func (s *Subscriber) Stop() {
	s.stopped = true
	close(s.goalEvents)
	close(s.cardEvents)
	close(s.matchStatusEvents)
}

func (s *Subscriber) GoalEvents() <-chan *GoalEvent {
	return s.goalEvents
}

func (s *Subscriber) CardEvents() <-chan *CardEvent {
	return s.cardEvents
}

func (s *Subscriber) MatchStatusEvents() <-chan *MatchStatusEvent {
	return s.matchStatusEvents
}
