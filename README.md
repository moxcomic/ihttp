## ihttp
A chained http call library

### Quick Start

```go
New().
	WithUrl("https://httpbin.org/get").
	Get().
	WithError(func(err error) { panic(err) }).
	ToString()
```

