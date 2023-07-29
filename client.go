package ihttp

import (
	"net/http"
)

type Client struct {
	*http.Client
}

func (self *IHttp) WithHijackRedirect() (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.client.CheckRedirect = func(req *http.Request, via []*http.Request) error {
		return http.ErrUseLastResponse
	}

	return
}

func (self *IHttp) do() {
	self.response.Response, self.err = self.client.Do(self.request.Request)
}

func (self *IHttp) Get() (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.Method = MethodGet
	self.do()

	return
}

func (self *IHttp) Post() (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.Method = MethodPost
	self.do()

	return
}

func (self *IHttp) Put() (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.Method = MethodPut
	self.do()

	return
}

func (self *IHttp) Delete() (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.Method = MethodDelete
	self.do()

	return
}

func (self *IHttp) Patch() (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.Method = MethodPatch
	self.do()

	return
}

func (self *IHttp) Head() (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.Method = MethodHead
	self.do()

	return
}

func (self *IHttp) Options() (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.request.Method = MethodOptions
	self.do()

	return
}
