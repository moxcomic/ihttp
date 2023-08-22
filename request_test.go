package ihttp

import (
	"fmt"
	"net/url"
	"testing"
	"time"
)

func TestError(t *testing.T) {
	s, err := New().
		Get().
		ToString()

	if err != nil {
		panic(err)
	}

	fmt.Println(s)
}

func TestGetString(t *testing.T) {
	fmt.Println(
		New().
			WithUrl("https://httpbin.org/get").
			Get().
			ToString(),
	)
}

func TestGetJson(t *testing.T) {
	v, err := New().
		WithUrl("https://httpbin.org/get").
		Get().
		ToJson()

	fmt.Println(v)
	fmt.Println(err)
	fmt.Println(v == nil)
	fmt.Println(v.GetString("origin"))
}

func TestGetGson(t *testing.T) {
	v, err := New().
		WithUrl("https://httpbin.org/get").
		Get().
		ToGson()

	fmt.Println(v)
	fmt.Println(err)
	fmt.Println(v.Nil())
	fmt.Println(v.Get("origin").Str())
}

func TestGetStruct(t *testing.T) {
	resp := struct {
		Origin string `json:"origin"`
	}{}

	err := New().
		WithUrl("https://httpbin.org/get").
		Get().
		ToJsonStruct(&resp)

	if err != nil {
		panic(err)
	}

	fmt.Println(resp.Origin)
}

func TestGet301(t *testing.T) {
	fmt.Println(
		New().
			WithUrl("https://b23.tv/26kpzlf").
			WithHijackRedirect().
			Head().
			ToLocation(),
	)
}

func TestProxy(t *testing.T) {
	fmt.Println(
		New().
			WithUrl("https://httpbin.org/get").
			WithLocalHttpProxy(7890).
			Get().
			ToString(),
	)
}

func TestTimeout(t *testing.T) {
	fmt.Println(
		New().
			WithUrl("https://httpbin.org/get").
			WithTimeout(time.Millisecond).
			Get().
			ToString(),
	)
}

func TestQuery(t *testing.T) {
	j1, err := New().
		WithUrl("https://httpbin.org/get").
		WithQuery(url.Values{"a": {"b"}}).
		Get().
		ToJson()

	if err != nil {
		panic(err)
	}

	fmt.Println(j1.GetString("url"))

	<-time.After(time.Second * 2)

	j2, err := New().
		WithUrl("https://httpbin.org/get").
		WithAddQuery("a", "b").
		Get().
		ToJson()

	if err != nil {
		panic(err)
	}

	fmt.Println(j2.GetString("url"))

	<-time.After(time.Second * 2)

	j3, err := New().
		WithUrl("https://httpbin.org/get").
		WithAddQuerys(map[string]any{"a": "b", "c": "d"}).
		Get().
		ToJson()

	if err != nil {
		panic(err)
	}

	fmt.Println(j3.GetString("url"))
}

func TestHeader(t *testing.T) {
	fmt.Println(
		New().
			WithUrl("https://httpbin.org/get").
			WithHeader("My-Header", "Header Value").
			Get().
			ToString(),
	)
}

func TestBody(t *testing.T) {
	fmt.Println(
		New().
			WithUrl("https://httpbin.org/post").
			WithHeader("Content-Type", "application/json").
			WithBody([]byte(`{"a":"b"}`)).
			Post().
			ToString(),
	)
}
