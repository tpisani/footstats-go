![Footstats Gopher](footstats-gopher.png)

# footstats-go

A Footstats SDK for [Go](https://golang.org/)

**WARNING: This package is still experimental. Future changes in public API may occur.**


## Usage

```go

import (
	"log"

	"github.com/tpisani/footstats-go"
)

func main() {
	c := footstats.NewClient("<footstats-baseurl>", "<user>", "<password>")

	matches, err := footstats.FilterMatches(c, func(m *footstats.Match) bool {
		return m.HasLiveCoverage
	})
	if err != nil {
		log.Fatal("Error while fetching matches:", err)
	}

	s := footstats.NewSubscriber(c, matches)
	s.Start()

	for {
		select {
		case e := <- s.GoalEvents():
			log.Printf("Goal: %+v\n", e.Goal)
		case e := <- s.CardEvents():
			log.Printf("Card: %+v\n", e.Card)
		case e := <- s.MatchStatusEvents():
			log.Printf("Match status: %+v\n", e)
		}
	}
}

```
