package ihttp

import (
	"fmt"
	"testing"
)

func TestError(t *testing.T) {
	var (
		e error
		s string
	)

	s = New().Get().WithError(func(err error) { e = err }).ToString()

	if e != nil {
		panic(e)
	}

	fmt.Println(s)
}

func TestGetString(t *testing.T) {
	fmt.Println(New().WithUrl("https://httpbin.org/get").Get().WithError(func(err error) { panic(err) }).ToString())
}

func TestGetJson(t *testing.T) {
	v := New().WithUrl("https://httpbin.org/get").Get().WithError(func(err error) { panic(err) }).ToJson()
	fmt.Println(v)
	fmt.Println(v == nil)
	fmt.Println(v.GetString("origin"))
}

func TestGetStruct(t *testing.T) {
	resp := struct {
		Origin string `json:"origin"`
	}{}

	New().WithUrl("https://httpbin.org/get").Get().WithError(func(err error) { panic(err) }).ToJsonStruct(&resp)

	fmt.Println(resp.Origin)
}

func TestGet301(t *testing.T) {
	fmt.Println(New().WithUrl("https://b23.tv/26kpzlf").WithHijackRedirect().WithError(func(err error) { panic(err) }).Head().ToLocation())
}

func TestProxy(t *testing.T) {
	fmt.Println(New().WithUrl("https://httpbin.org/get").WithLocalHttpProxy(7890).Get().WithError(func(err error) { panic(err) }).ToString())
}
