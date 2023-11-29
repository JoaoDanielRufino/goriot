package internal

import (
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

const riotHeaderApiKey = "X-Riot-Token"

type HttpClient struct {
	client  *http.Client
	baseURL string
}

type RequestModifier func(*http.Request)

func NewDefaultHttpClient(baseURL string) *HttpClient {
	return &HttpClient{
		client:  http.DefaultClient,
		baseURL: baseURL,
	}
}

func NewInsecureHttpClient(baseURL string) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		baseURL: baseURL,
	}
}

func NewSecureHttpClient(baseURL string, certificate []byte) *HttpClient {
	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(certificate)

	return &HttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: certPool,
				},
			},
		},
		baseURL: baseURL,
	}
}

func WithApiKey(key string) RequestModifier {
	return func(r *http.Request) {
		r.Header.Add(riotHeaderApiKey, key)
	}
}

func (c *HttpClient) Do(method, endpoint string, body io.Reader, modifiers ...RequestModifier) (*http.Response, error) {
	req, err := c.newRequest(method, endpoint, body, modifiers...)
	if err != nil {
		return nil, err
	}

	res, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	if res.StatusCode < 200 || res.StatusCode > 299 {
		return nil, c.handleErrorStatusCode(res)
	}

	return res, nil
}

func (c *HttpClient) newRequest(method, endpoint string, body io.Reader, modifiers ...RequestModifier) (*http.Request, error) {
	url := c.baseURL + endpoint
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}

	for _, modifier := range modifiers {
		modifier(req)
	}

	return req, nil
}

func (c *HttpClient) handleErrorStatusCode(res *http.Response) error {
	body := struct {
		Status struct {
			Message    string `json:"message"`
			StatusCode uint   `json:"status_code"`
		} `json:"status"`
	}{}

	defer res.Body.Close()

	err := json.NewDecoder(res.Body).Decode(&body)
	if err != nil {
		return fmt.Errorf("Error: %s", res.Status)
	}

	if body.Status.Message == "" {
		return fmt.Errorf("Error: %s", res.Status)
	}

	return RequestError{
		Message:    body.Status.Message,
		StatusCode: body.Status.StatusCode,
	}
}
