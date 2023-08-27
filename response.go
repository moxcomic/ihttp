package ihttp

import (
	"bytes"
	"net/http"
	"strings"

	"github.com/spf13/viper"
	"github.com/ysmood/gson"
)

type Response struct {
	*http.Response
}

func (self *IHttp) ToHeader() (http.Header, error) {
	if self.err != nil {
		return nil, self.err
	}

	return self.response.Header, nil
}

func (self *IHttp) ToLocation() (string, error) {
	if self.err != nil {
		return "", self.err
	}

	return self.response.Header.Get("Location"), nil
}

func (self *IHttp) ToBytes() ([]byte, error) {
	if self.err != nil {
		return nil, self.err
	}

	return self.responseData, nil
}

func (self *IHttp) ToString() (string, error) {
	if self.err != nil {
		return "", self.err
	}

	return string(self.responseData), self.err
}

func (self *IHttp) ToJson() (*viper.Viper, error) {
	if self.err != nil {
		return nil, self.err
	}

	v := viper.New()
	v.SetConfigType("json")
	self.err = v.ReadConfig(bytes.NewReader(self.responseData))

	return v, self.err
}

func (self *IHttp) ToGson() (gson.JSON, error) {
	if self.err != nil {
		return gson.New(""), self.err
	}

	str, err := self.ToString()
	if err != nil {
		return gson.New(""), err
	}

	return gson.NewFrom(str), nil
}

func (self *IHttp) ToJsonStruct(value any) error {
	j, err := self.ToJson()
	if err != nil {
		return err
	}

	return j.Unmarshal(value)
}

func (self *IHttp) ToReader() (*strings.Reader, error) {
	if self.err != nil {
		return nil, self.err
	}

	return strings.NewReader(string(self.responseData)), self.err
}
