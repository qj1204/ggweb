package main

import (
	"fmt"
	"log"
	"net/http"
)

// 源码：ListenAndServe(addr string, handler Handler)
// Handler是一个接口，包含一个ServeHTTP方法，这个方法接收一个http.ResponseWriter和一个*http.Request，需要自己实现该接口

type Engine struct{}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func main() {
	engine := &Engine{}
	log.Fatal(http.ListenAndServe(":9999", engine))
}
