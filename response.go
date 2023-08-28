package ihttp

import (
	"bytes"
	"io"
	"net/http"

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

func (self *IHttp) ToReader() (*bytes.Reader, error) {
	if self.err != nil {
		return nil, self.err
	}

	return bytes.NewReader(self.responseData), self.err
}

func (self *IHttp) ToNopCloser() (io.ReadCloser, error) {
	if self.err != nil {
		return nil, self.err
	}

	var reader *bytes.Reader
	reader, self.err = self.ToReader()
	if self.err != nil {
		return nil, self.err
	}

	return io.NopCloser(reader), self.err
}
