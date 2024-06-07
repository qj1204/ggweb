package main

import (
	"ggin"
	"net/http"
)

func main() {
	r := ggin.New()
	r.GET("/", func(c *ggin.Context) {
		c.HTML(http.StatusOK, "<h1>Hello ggin</h1>")
	})
	r.GET("/hello", func(c *ggin.Context) {
		// expect /hello?name=gginktutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})
	r.GET("/hello/:name", func(c *ggin.Context) {
		// expect /hello/geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})

	r.GET("/assets/*filepath", func(c *ggin.Context) {
		c.JSON(http.StatusOK, ggin.H{"filepath": c.Param("filepath")})
	})

	r.Run(":9999")

}
