package ggin

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// H 这个就是gin.H的那个H
type H map[string]any

type Context struct {
	// origin objects
	Writer http.ResponseWriter
	Req    *http.Request
	// request info
	Path   string
	Method string
	// response info
	StatusCode int
}

func newContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer: w,
		Req:    req,
		Path:   req.URL.Path,
		Method: req.Method,
	}
}

// PostForm 获取form参数
func (c *Context) PostForm(key string) string {
	return c.Req.FormValue(key)
}

// Query 获取查询参数
func (c *Context) Query(key string) string {
	return c.Req.URL.Query().Get(key)
}

// Status 设置状态码
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader 设置响应头
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String 构造String响应
func (c *Context) String(code int, format string, values ...any) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)

	if len(values) > 0 {
		c.Writer.Write([]byte(fmt.Sprintf(format, values...)))
		return
	}
	c.Writer.Write([]byte(format))
}

// JSON 构造JSON响应
func (c *Context) JSON(code int, obj any) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	jsonBytes, err := json.Marshal(obj)
	if err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
	c.Writer.Write(jsonBytes)
}

// Data 构造字节流响应
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML 构造HTML响应
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
