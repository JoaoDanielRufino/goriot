package live

import (
	"io"

	"github.com/JoaoDanielRufino/goriot/internal"
)

type LiveClientData struct {
	httpClient *internal.HttpClient
}

func NewLiveClientData(httpClient *internal.HttpClient) *LiveClientData {
	return &LiveClientData{
		httpClient: httpClient,
	}
}

func (c *LiveClientData) AllGameData() (string, error) {
	resp, err := c.httpClient.Do("GET", allGameDataEndpoint, nil)
	if err != nil {
		return "", err
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}
