package footstats

import (
	"encoding/json"
	"strconv"
)

type FeedAction int

const (
	MatchStart FeedAction = iota
	MatchEnd
	GoalAction
	OwnGoalAction
)

type FeedItem struct {
	ID          int
	MatchID     int
	TeamID      int
	PlayerID    int
	PlayerName  string
	Action      string
	Description string
	Period      MatchPeriod
	Minute      int
}

type Feed []FeedItem

type footstatsFeedResponse struct {
	Feed Feed `json:"Narracoes"`
}

type foostatsFeedItem struct {
	ID          string `json:"Id"`
	MatchID     string `json:"IdPartida"`
	TeamID      string `json:"IdEquipe"`
	PlayerID    string `json:"IdJogador"`
	PlayerName  string `json:"NomeJogador"`
	Action      string `json:"AcaoImportante"`
	Description string `json:"Descricao"`
	Period      string `json:"Periodo"`
	Minute      string `json:"Momento"`
}

func (f *FeedItem) UnmarshalJSON(data []byte) error {
	var o foostatsFeedItem

	err := json.Unmarshal(data, &o)
	if err != nil {
		return err
	}

	id, _ := strconv.Atoi(o.ID)
	matchID, _ := strconv.Atoi(o.MatchID)
	teamID, _ := strconv.Atoi(o.TeamID)
	playerID, _ := strconv.Atoi(o.PlayerID)
	period := matchPeriodFromString(o.Period)
	minute, _ := strconv.Atoi(o.Minute)

	f.ID = id
	f.MatchID = matchID
	f.TeamID = teamID
	f.PlayerID = playerID
	f.PlayerName = o.PlayerName
	f.Action = o.Action
	f.Description = o.Description
	f.Period = period
	f.Minute = minute

	return nil
}
