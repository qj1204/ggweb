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
	r.GET("/hello2", func(c *ggin.Context) {
		c.String(http.StatusOK, "hello world")
	})

	r.POST("/login", func(c *ggin.Context) {
		c.JSON(http.StatusOK, ggin.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")

}
