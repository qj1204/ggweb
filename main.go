package main

import (
	"ggin"
	"net/http"
)

func main() {
	r := ggin.New()
	r.GET("/index", func(c *ggin.Context) {
		c.HTML(http.StatusOK, "<h1>Index Page</h1>")
	})
	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *ggin.Context) {
			c.HTML(http.StatusOK, "<h1>Hello ggin</h1>")
		})

		v1.GET("/hello", func(c *ggin.Context) {
			// expect /hello?name=gginktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
		})
	}
	v2 := r.Group("/v2")
	{
		v2.GET("/hello/:name", func(c *ggin.Context) {
			// expect /hello/gginktutu
			c.String(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
		})
		v2.POST("/login", func(c *ggin.Context) {
			c.JSON(http.StatusOK, ggin.H{
				"username": c.PostForm("username"),
				"password": c.PostForm("password"),
			})
		})

	}

	r.Run(":9999")
}
