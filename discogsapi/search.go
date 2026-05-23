package discogsapi

import (
	"encoding/json"
	"net/http"
	"net/url"
)

type SearchResponse struct {
	Results []SearchResult `json:"results"`
}

type SearchResult struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  string `json:"year"`
}

func (c *Client) AlbumSearch(query string) ([]SearchResult, error) {
	normalized := normalize(query)

	params := url.Values{}
	params.Set("release_title", normalized)
	params.Set("type", "master")
	params.Set("per_page", "10")

	endpoint := c.baseURL + "/database/search?" + params.Encode()

	req, err := http.NewRequest(http.MethodGet, endpoint, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", "discography-cli/0.1")
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	var searchResp SearchResponse

	err = json.NewDecoder(resp.Body).Decode(&searchResp)
	if err != nil {
		return nil, err
	}

	return searchResp.Results, nil
}
