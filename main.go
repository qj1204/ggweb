package main

import (
	"fmt"
	"ggin"
	"net/http"
	"time"
)

type student struct {
	Name string
	Age  int8
}

func FormatAsDate(t time.Time) string {
	year, month, day := t.Date()
	return fmt.Sprintf("%d-%02d-%02d", year, month, day)
}

func main() {
	r := ggin.Default()
	r.GET("/", func(c *ggin.Context) {
		c.String(http.StatusOK, "Hello xiaoxin\n")
	})
	// index out of range for testing Recovery()
	r.GET("/panic", func(c *ggin.Context) {
		names := []string{"xiaoxin"}
		c.String(http.StatusOK, names[100])
	})

	r.Run(":9999")
}
