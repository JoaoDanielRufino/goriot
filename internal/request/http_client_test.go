package request

import (
	"bytes"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

const baseURL = "http://BASE_URL"

func TestNewDefaultHttpClient(t *testing.T) {
	expected := &HttpClient{
		client:  http.DefaultClient,
		baseURL: baseURL,
	}

	got := NewDefaultHttpClient(baseURL)

	assert.Equal(t, expected, got)
}

func TestNewInsecureHttpClient(t *testing.T) {
	expected := &HttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					InsecureSkipVerify: true,
				},
			},
		},
		baseURL: baseURL,
	}

	got := NewInsecureHttpClient(baseURL)

	assert.Equal(t, expected, got)
}

func TestNewSecureHttpClient(t *testing.T) {
	certificate := []byte{}

	certPool := x509.NewCertPool()
	certPool.AppendCertsFromPEM(certificate)

	expected := &HttpClient{
		client: &http.Client{
			Transport: &http.Transport{
				TLSClientConfig: &tls.Config{
					RootCAs: certPool,
				},
			},
		},
		baseURL: baseURL,
	}

	got := NewSecureHttpClient(baseURL, certificate)

	assert.Equal(t, expected, got)
}

type RoundTripperMock struct {
	Response *http.Response
}

func (r *RoundTripperMock) RoundTrip(req *http.Request) (*http.Response, error) {
	return r.Response, nil
}

func TestDo(t *testing.T) {
	tests := []struct {
		name      string
		client    *HttpClient
		modifiers []RequestModifier
		want      *http.Response
		err       error
	}{
		{
			name:   "should successfully make a http request without modifiers",
			client: newMockedHttpClient(baseURL, http.StatusAccepted, "OK"),
			want:   newHttpResponse(http.StatusAccepted, "OK"),
			err:    nil,
		},
		{
			name:      "should successfully make a http request with modifiers",
			client:    newMockedHttpClient(baseURL, http.StatusAccepted, "OK"),
			modifiers: []RequestModifier{WithApiKey("API_KEY")},
			want:      newHttpResponse(http.StatusAccepted, "OK"),
			err:       nil,
		},
		{
			name:   "should return generic error when response body decode fails",
			client: newMockedHttpClient(baseURL, http.StatusNotFound, ""),
			want:   nil,
			err:    fmt.Errorf(genericErrorMessage, "404"),
		},
		{
			name:   "should return RequestError when response is with error status code and response body matches expected struct",
			client: newMockedHttpClient(baseURL, http.StatusNotFound, newErrorBody("not found", 404)),
			want:   nil,
			err:    RequestError{Message: "not found", StatusCode: 404},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.client.Do("GET", "ENDPOINT", nil, tt.modifiers...)

			assert.Equal(t, tt.err, err)
			assert.Equal(t, tt.want, got)
		})
	}
}

func newMockedHttpClient(baseURL string, statusCode int, body interface{}) *HttpClient {
	return &HttpClient{
		client: &http.Client{
			Transport: &RoundTripperMock{
				Response: newHttpResponse(statusCode, body),
			},
		},
		baseURL: baseURL,
	}
}

func newHttpResponse(statusCode int, body interface{}) *http.Response {
	data, _ := json.Marshal(body)
	reader := bytes.NewReader(data)

	return &http.Response{
		Status:     fmt.Sprint(statusCode),
		StatusCode: statusCode,
		Body:       io.NopCloser(reader),
	}
}

func newErrorBody(message string, statusCode uint) interface{} {
	type status struct {
		Message    string `json:"message"`
		StatusCode uint   `json:"status_code"`
	}

	type body struct {
		Status status
	}

	return body{
		Status: status{
			Message:    message,
			StatusCode: statusCode,
		},
	}
}
