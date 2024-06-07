package main

import (
	"ggin"
	"log"
	"net/http"
	"time"
)

func onlyForV2() ggin.HandlerFunc {
	return func(c *ggin.Context) {
		// Start timer
		t := time.Now()
		// if a server error occurred
		c.Fail(500, "Internal Server Error")
		// Calculate resolution time
		log.Printf("[%d] %s in %v for group v2", c.StatusCode, c.Req.RequestURI, time.Since(t))
	}
}

func main() {
	r := ggin.New()
	r.Use(ggin.Logger()) // global midlleware
	r.GET("/", func(c *ggin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ggin</h1>")
	})

	v2 := r.Group("/v2")
	v2.Use(onlyForV2()) // v2 group middleware
	{
		v2.GET("/hello/:name", func(c *ggin.Context) {
			// expect /hello/gginktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
	}

	r.Run(":9999")
}
