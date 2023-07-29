package ihttp

import (
	"io"
	"net/http"

	"github.com/spf13/viper"
	"github.com/ysmood/gson"
)

type Response struct {
	*http.Response
}

func (self *IHttp) ToHeader() (resp http.Header) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	resp = self.response.Header

	return
}

func (self *IHttp) ToLocation() (resp string) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	resp = self.response.Header.Get("Location")

	return
}

func (self *IHttp) ToBytes() (resp []byte) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	resp, self.err = io.ReadAll(self.response.Body)

	return
}

func (self *IHttp) ToString() (resp string) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	var data []byte

	data, self.err = io.ReadAll(self.response.Body)

	return string(data)
}

func (self *IHttp) ToJson() (resp *viper.Viper) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	resp = viper.New()
	resp.SetConfigType("json")
	self.err = resp.ReadConfig(self.response.Body)

	return
}

func (self *IHttp) ToGson() (resp gson.JSON) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	data := self.ToString()

	resp = gson.NewFrom(data)

	return
}

func (self *IHttp) ToJsonStruct(value any) {
	defer self.doErrorHandler()

	self.err = self.ToJson().Unmarshal(value)
}
