package footstats

import (
	"sync"
	"time"
)

type matchFilter func(*Match) bool

func filterMatches(wg *sync.WaitGroup, client *Client, c *Championship, todaysMatches *[]*Match, filter matchFilter) {
	defer wg.Done()

	matches, err := client.Matches(c.ID)
	if err != nil {
		return
	}

	for _, m := range matches {
		if filter(m) {
			*todaysMatches = append(*todaysMatches, m)
		}
	}

}

func FilterMatches(client *Client, filter matchFilter) ([]*Match, error) {
	var todaysMatches []*Match

	championships, err := client.Championships()
	if err != nil {
		return todaysMatches, err
	}

	wg := &sync.WaitGroup{}

	for _, championship := range championships {
		wg.Add(1)
		go filterMatches(wg, client, championship, &todaysMatches, filter)
	}

	wg.Wait()

	return todaysMatches, nil
}

func TodaysMatches(client *Client) ([]*Match, error) {
	curYear, curMonth, curDay := time.Now().Date()

	return FilterMatches(client, func(m *Match) bool {
		year, month, day := m.ScheduledTo.Date()
		return year == curYear && month == curMonth && day == curDay
	})
}
