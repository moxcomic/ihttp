package ihttp

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"net/url"

	"golang.org/x/net/proxy"
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

func (self *IHttp) WithLocalHttpsProxy(p int) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	self.client.Transport = &http.Transport{
		Proxy: func(r *http.Request) (*url.URL, error) {
			return url.Parse(fmt.Sprintf("https://localhost:%d", p))
		},
	}

	return
}

func (self *IHttp) WithLocalSocks5Proxy(p int) (this *IHttp) {
	this = self

	if self.err != nil {
		return
	}

	var dialer proxy.Dialer
	dialer, self.err = proxy.SOCKS5("tcp", fmt.Sprintf("127.0.0.1:%d", p), nil, nil)
	if self.err != nil {
		return
	}

	self.client.Transport = &http.Transport{
		DialContext: func(ctx context.Context, network, addr string) (net.Conn, error) {
			return dialer.Dial(network, addr)
		},
	}

	return
}
