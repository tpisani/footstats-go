package footstats

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"regexp"
)

type Client struct {
	baseURL  string
	user     string
	password string
}

func (c *Client) buildURL(endpoint string) string {
	return fmt.Sprintf("%s%s?usuario=%s&senha=%s",
		c.baseURL, endpoint, c.user, c.password)
}

func (c *Client) makeRequest(endpoint string) ([]byte, error) {
	url := c.buildURL(endpoint)

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
	tagExp := regexp.MustCompile("</?.*>")

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
	data, err := makeRequest("ListaCampeonatos")
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
