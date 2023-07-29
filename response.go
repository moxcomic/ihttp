package ihttp

import (
	"bytes"
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

	resp = self.responseData

	return
}

func (self *IHttp) ToString() (resp string) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	resp = string(self.responseData)

	return
}

func (self *IHttp) ToJson() (resp *viper.Viper) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	resp = viper.New()
	resp.SetConfigType("json")
	self.err = resp.ReadConfig(bytes.NewReader(self.responseData))

	return
}

func (self *IHttp) ToGson() (resp gson.JSON) {
	defer self.doErrorHandler()

	if self.err != nil {
		return
	}

	resp = gson.NewFrom(self.ToString())

	return
}

func (self *IHttp) ToJsonStruct(value any) {
	defer self.doErrorHandler()

	self.err = self.ToJson().Unmarshal(value)
}
