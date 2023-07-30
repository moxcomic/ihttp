package ihttp

import (
	"net/http"
	"net/url"
	"sync"
	"time"
)

type IHttp struct {
	err error

	client       *Client
	request      *Request
	response     *Response
	responseData []byte

	onceError    sync.Once
	errorHandler func(error)
}

func New() *IHttp {
	return &IHttp{
		client: &Client{Client: &http.Client{
			Timeout: time.Second * 60,
		}},
		request: &Request{
			Request: &http.Request{
				Header: make(http.Header),
			},
			queryParams: url.Values{},
		},
		response: &Response{Response: &http.Response{}},
	}
}
