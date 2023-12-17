package mock

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/JoaoDanielRufino/goriot/internal/request"
)

type RequesterMock struct {
	response *http.Response
	err      error
}

func JsonResponseMock(target interface{}, err error) RequesterMock {
	data, _ := json.Marshal(target)
	reader := bytes.NewReader(data)

	return RequesterMock{
		response: &http.Response{
			Body: io.NopCloser(reader),
		},
		err: err,
	}
}

func (r RequesterMock) Do(method, endpoint string, body io.Reader, modifiers ...request.RequestModifier) (*http.Response, error) {
	return r.response, r.err
}
