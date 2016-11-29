package footstats

import (
	"sync"
)

type Subscriber struct {
	initialized bool

	client  *Client
	matches []*Match

	goalMutex  sync.Mutex
	goalIDs    []int
	goalEvents chan *GoalEvent
}

func NewSubscriber(c *Client, matches []*Match) *Subscriber {
	s := &Subscriber{
		client:  c,
		matches: matches,

		goalEvents: make(chan *GoalEvent),
	}

	go s.startPolling()

	return s
}

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

func (s *Subscriber) checkNewGoals(wg *sync.WaitGroup, m *Match, goals []*Goal) {
	defer wg.Done()

	for _, g := range goals {
		if s.addGoal(g) && s.initialized {
			s.goalEvents <- &GoalEvent{
				Match: m,
				Goal:  g,
			}
		}
	}
}

func (s *Subscriber) poll(wg *sync.WaitGroup, m *Match) {
	defer wg.Done()

	data, err := s.client.MatchStats(m.ID)
	if err != nil {
		return
	}

	wg.Add(1)
	go s.checkNewGoals(wg, m, data.Goals)
}

func (s *Subscriber) startPolling() {
	wg := &sync.WaitGroup{}

	for {
		for _, m := range s.matches {
			wg.Add(1)
			go s.poll(wg, m)
		}

		wg.Wait()
		s.initialized = true
	}
}

func (s *Subscriber) GoalEvents() <-chan *GoalEvent {
	return s.goalEvents
}
