package serviceac

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/hashicorp/go-retryablehttp"
	"io"
	"net/http"
	"time"
)

const (
	baseURL = `https://api.appcenter.ms`
)

func (rt roundTripper) RoundTrip(req *http.Request) (*http.Response, error) {
	req.Header.Add(
		"x-api-token", rt.token,
	)
	req.Header.Add(
		"content-type", "application/json; charset=utf-8",
	)
	return http.DefaultTransport.RoundTrip(req)
}

type roundTripper struct {
	token string
}

type Client struct {
	httpClient *retryablehttp.Client
}

func NewClient(token string) Client {
	retClient := retryablehttp.NewClient()

	retClient.RetryMax = 5
	retClient.RetryWaitMin = 5 * time.Second
	retClient.RetryWaitMax = 10 * time.Second

	retClient.HTTPClient.Transport = &roundTripper{
		token: token,
	}

	return Client{httpClient: retClient}
}

func (c Client) jsonRequest(method, url string, body []byte, response interface{}) (int, error) {
	var reader io.Reader

	if body != nil {
		reader = bytes.NewReader(body)
	}

	req, err := retryablehttp.NewRequest(method, url, reader)

	if err != nil {
		return -1, err
	}

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return -1, err
	}

	defer func() {
		if err := resp.Body.Close(); err != nil {
		}
	}()

	if response != nil {
		rb, err := io.ReadAll(resp.Body)
		if err != nil {
			return -1, err
		}

		if err := json.Unmarshal(rb, response); err != nil {
			return resp.StatusCode, fmt.Errorf("error: %s, response: %s", err, string(rb))
		}
	}

	return resp.StatusCode, nil
}
