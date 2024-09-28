package main

import (
	"go-example/patterns/structural/decorator"
	"net/http"
)

func main() {
	http.HandleFunc("/v1/hello", decorator.Handler(
		decorator.Hello,
		decorator.WithServerHeader,
		decorator.WithBasicAuth,
		decorator.WithDebugLog,
	))
}
