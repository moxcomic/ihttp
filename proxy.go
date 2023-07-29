package ihttp

import (
	"fmt"
	"net/http"
	"net/url"
)

func (self *IHttp) WithProxy(p func(r *http.Request) (*url.URL, error)) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.client.Transport = &http.Transport{
		Proxy: p,
	}

	return
}

func (self *IHttp) WithLocalHttpProxy(p int) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.client.Transport = &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			return url.Parse(fmt.Sprintf("http://localhost:%d", p))
		},
	}

	return
}

func (self *IHttp) WithLocalSocks5Proxy(p int) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.client.Transport = &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			return url.Parse(fmt.Sprintf("socks5://localhost:%d", p))
		},
	}

	return
}
