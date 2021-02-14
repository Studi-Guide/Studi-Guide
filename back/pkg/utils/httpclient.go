package utils

import "net/http"

//HttpClient to wrap the http request
type HttpClient interface {
	Do(req *http.Request) (*http.Response, error)
}
