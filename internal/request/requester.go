package request

import (
	"io"
	"net/http"
)

type Requester interface {
	Do(method, endpoint string, body io.Reader, modifiers ...RequestModifier) (*http.Response, error)
}
