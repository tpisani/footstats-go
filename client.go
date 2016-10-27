package footstats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"regexp"
	"strconv"
)

var tagExp = regexp.MustCompile("<.*?>")

type Client struct {
	baseURL  string
	user     string
	password string
}

func NewClient(baseURL, user, password string) *Client {
	return &Client{
		baseURL:  baseURL,
		user:     user,
		password: password,
	}
}

func (c *Client) buildURL(endpoint string, params *url.Values) string {
	u := fmt.Sprintf("%s/%s?usuario=%s&senha=%s",
		c.baseURL, endpoint, c.user, c.password)

	if params != nil {
		u = fmt.Sprintf("%s&%s", u, params.Encode())
	}

	return u
}

func (c *Client) makeRequest(endpoint string, params *url.Values) ([]byte, error) {
	u := c.buildURL(endpoint, params)

	resp, err := http.Get(u)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	data = tagExp.ReplaceAll(data, []byte(""))

	return data, nil
}

func (c *Client) Championships() ([]*Championship, error) {
	data, err := c.makeRequest("ListaCampeonatos", nil)
	if err != nil {
		return nil, err
	}

	var wrapper championshipWrapper
	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Championships, nil
}

func (c *Client) Matches(championshipId int) ([]*Match, error) {
	params := &url.Values{}
	params.Set("campeonato", strconv.Itoa(championshipId))

	data, err := c.makeRequest("ListaPartidas", params)
	if err != nil {
		return nil, err
	}

	var wrapper matchWrapper
	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	return wrapper.Matches.Match, nil
}

func (c *Client) Entities(championshipId int) (*Entities, error) {
	params := &url.Values{}
	params.Set("campeonato", strconv.Itoa(championshipId))

	data, err := c.makeRequest("ListaEntidades", params)
	if err != nil {
		return nil, err
	}

	var wrapper entitiesWrapper
	err = json.Unmarshal(data, &wrapper)
	if err != nil {
		return nil, err
	}

	return &Entities{
		Teams:    wrapper.Teams.Team,
		Players:  wrapper.Players.Player,
		Coaches:  wrapper.Coaches.Coach,
		Referees: wrapper.Referees.Referee,
		Stadiums: wrapper.Stadiums.Stadium,
	}, nil
}

func (c *Client) MatchData(matchId int) (*MatchData, error) {
	params := &url.Values{}
	params.Set("idpartida", strconv.Itoa(matchId))

	data, err := c.makeRequest("AoVivo", params)
	if err != nil {
		return nil, err
	}

	var events *MatchData
	err = json.Unmarshal(data, &events)
	if err != nil {
		return nil, err
	}

	return events, nil
}
