package footstats

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strconv"
	"sync"
)

type matchFilter func(match Match) bool

type nopReadCloser struct{}

func (r nopReadCloser) Read(p []byte) (n int, err error) {
	return 0, nil
}

func (n nopReadCloser) Close() error {
	return nil
}

type Client struct {
	baseURL string
	token   string
}

func NewClient(token string) *Client {
	return &Client{
		baseURL: "http://apifutebol.footstats.com.br",
		token:   token,
	}
}

func (c Client) buildURL(endpoint string, params *url.Values) string {
	var b bytes.Buffer

	fmt.Fprintf(&b, "%s/%s?token=%s", c.baseURL, endpoint, c.token)

	if params != nil {
		fmt.Fprintf(&b, "&%s", params.Encode())
	}

	return b.String()
}

func (c Client) makeRequest(endpoint string, params *url.Values) (io.ReadCloser, error) {
	u := c.buildURL(endpoint, params)
	resp, err := http.Get(u)
	if err != nil {
		return nopReadCloser{}, err
	}

	statusFamily := resp.StatusCode / 100

	if statusFamily == 4 {
		err = errors.New("client error: " + resp.Status)
	} else if statusFamily == 5 {
		err = errors.New("internal server error: " + resp.Status)
	}

	return resp.Body, err
}

func (c Client) Championships() ([]Championship, error) {
	var championships []Championship

	r, err := c.makeRequest("V2/api/Campeonato/ListarCampeonatos", nil)
	if err != nil {
		return championships, err
	}
	defer r.Close()

	err = json.NewDecoder(r).Decode(&championships)

	return championships, err
}

func (c Client) TeamsByChampionship(championshipID int) ([]Team, error) {
	var teams []Team

	params := &url.Values{}
	params.Set("idcampeonato", strconv.Itoa(championshipID))

	r, err := c.makeRequest("V2/api/Equipe/EquipesCampeonato", params)
	if err != nil {
		return teams, err
	}
	defer r.Close()

	err = json.NewDecoder(r).Decode(&teams)

	return teams, err
}

func (c Client) PlayersByTeam(teamID int) ([]Player, error) {
	var players []Player

	params := &url.Values{}
	params.Set("idequipe", strconv.Itoa(teamID))
	r, err := c.makeRequest("V2/api/Jogador/JogadoresEquipe", params)
	if err != nil {
		return players, err
	}

	err = json.NewDecoder(r).Decode(&players)

	return players, err
}

func (c Client) MatchesByChampionship(championshipID int) ([]Match, error) {
	var matches []Match

	params := &url.Values{}
	params.Set("idcampeonato", strconv.Itoa(championshipID))

	r, err := c.makeRequest("V2/api/Partida/PartidasCampeonato", params)
	if err != nil {
		return matches, err
	}
	defer r.Close()

	err = json.NewDecoder(r).Decode(&matches)

	return matches, err
}

func (c Client) OnGoingMatches() ([]Match, error) {
	var matches []Match

	r, err := c.makeRequest("V2/api/Partida/PartidasAndamento", nil)
	if err != nil {
		return matches, err
	}
	defer r.Close()

	err = json.NewDecoder(r).Decode(&matches)

	return matches, err
}

func (c Client) filterMatches(wg *sync.WaitGroup, championshipID int, filteredMatches *[]Match, filter matchFilter) {
	defer wg.Done()

	matches, err := c.MatchesByChampionship(championshipID)
	if err != nil {
		fmt.Println(err)
		return
	}

	for _, m := range matches {
		if filter(m) {
			*filteredMatches = append(*filteredMatches, m)
		}
	}
}

func (c Client) FilterMatches(filter matchFilter) ([]Match, error) {
	var matches []Match

	championships, err := c.Championships()
	if err != nil {
		return matches, err
	}

	wg := &sync.WaitGroup{}

	for _, championship := range championships {
		wg.Add(1)
		go c.filterMatches(wg, championship.ID, &matches, filter)
	}

	wg.Wait()

	return matches, nil
}

func (c Client) MatchFeed(matchID int) (Feed, error) {
	var feedResponse footstatsFeedResponse

	params := &url.Values{}
	params.Set("idpartida", strconv.Itoa(matchID))
	r, err := c.makeRequest("V2/api/Partida/NarracaoMinMin", params)
	if err != nil {
		return Feed{}, err
	}
	defer r.Close()

	err = json.NewDecoder(r).Decode(&feedResponse)

	return feedResponse.Feed, err
}

func (c Client) MatchLineup(matchID int) (MatchLineup, error) {
	var lineup MatchLineup

	params := &url.Values{}
	params.Set("idpartida", strconv.Itoa(matchID))
	r, err := c.makeRequest("V2/api/Partida/Escalacao", params)
	if err != nil {
		return lineup, err
	}
	defer r.Close()

	err = json.NewDecoder(r).Decode(&lineup)

	return lineup, err
}
