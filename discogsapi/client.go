package search

import "net/http"

type Client struct {
	httpClient http.Client
	baseURL    string
}
