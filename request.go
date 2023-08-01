package ihttp

import (
	"fmt"
	"net/http"
	"net/url"
)

type Request struct {
	*http.Request
	queryParams url.Values
}

func (self *IHttp) WithUrl(u string) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	value, err := url.Parse(u)
	if err != nil {
		self.err = err
		return
	}

	self.request.URL = value

	return
}

func (self *IHttp) WithQuery(q url.Values) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	for k, v := range q {
		self.request.queryParams.Add(k, v[0])
	}

	return
}

func (self *IHttp) WithAddQuery(k, v string) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.queryParams.Add(k, v)

	return
}

func (self *IHttp) WithAddQuerys(q map[string]any) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	for k, v := range q {
		self.request.queryParams.Add(k, fmt.Sprintf("%v", v))
	}

	return
}

func (self *IHttp) WithHeader(k, v string) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.Header.Add(k, v)

	return
}

func (self *IHttp) WithHeaders(h map[string]string) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	for k, v := range h {
		self.request.Header.Add(k, v)
	}

	return
}

func (self *IHttp) WithCookie(v string) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.WithHeader("Cookie", v)

	return
}
