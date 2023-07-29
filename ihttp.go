package ihttp

import (
	"net/http"
	"sync"
)

type IHttp struct {
	err error

	client   *Client
	request  *Request
	response *Response

	onceError    sync.Once
	errorHandler func(error)
}

func New() *IHttp {
	return &IHttp{
		client:   &Client{Client: &http.Client{}},
		request:  &Request{Request: &http.Request{}},
		response: &Response{Response: &http.Response{}},
	}
}
