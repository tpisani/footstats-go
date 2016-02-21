package footstats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	gourl "net/url"
	"regexp"
	"strconv"
)

type Client struct {
	baseURL  string
	user     string
	password string
}

func (c *Client) buildURL(endpoint string, params *gourl.Values) string {
	url := fmt.Sprintf("%s%s?usuario=%s&senha=%s",
		c.baseURL, endpoint, c.user, c.password)

	if params != nil {
		url = fmt.Sprintf("%s&%s", url, params.Encode())
	}

	return url
}

func (c *Client) makeRequest(endpoint string, params *gourl.Values) ([]byte, error) {
	url := c.buildURL(endpoint, params)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	// TODO: Make this a constant somehow. It causes a build
	// error when put on const declarations:
	// const initializer regexp.MustCompile(...) is not a constant
	tagExp := regexp.MustCompile("<.*?>")

	data = tagExp.ReplaceAll(data, []byte(""))

	return data, nil
}

func NewClient(baseURL, user, password string) *Client {
	return &Client{
		baseURL:  baseURL,
		user:     user,
		password: password,
	}
}

func (c *Client) Championships() ([]*Championship, error) {
	data, err := c.makeRequest("ListaCampeonatos", nil)
	if err != nil {
		return nil, err
	}

	var footstatsData championshipData
	err = json.Unmarshal(data, &footstatsData)
	if err != nil {
		return nil, err
	}

	return footstatsData.championships(), nil
}

func (c *Client) Matches(championshipId int64) ([]*Match, error) {
	params := &gourl.Values{}
	params.Set("campeonato", strconv.FormatInt(championshipId, 10))

	data, err := c.makeRequest("ListaPartidas", params)
	if err != nil {
		return nil, err
	}

	var footstatsData matchData
	err = json.Unmarshal(data, &footstatsData)
	if err != nil {
		return nil, err
	}

	return footstatsData.matches(championshipId), nil
}
