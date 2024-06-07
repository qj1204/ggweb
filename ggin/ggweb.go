package ggin

import (
	"net/http"
)

// HandlerFunc 提供给用户，用来定义路由映射的处理方法
type HandlerFunc func(c *Context)

// Engine 实现了http.Handler接口
type Engine struct {
	router *router // 路由和处理函数映射表
}

// New 创建一个Engine实例
func New() *Engine {
	return &Engine{router: newRouter()}
}

// ServeHTTP 解析请求的路径，查找路由映射表，如果查到，就执行注册的处理方法。如果查不到，就返回 404 NOT FOUND
func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	c := newContext(w, req)
	engine.router.handle(c)
}

// addRoute 将请求方式和请求路径以及处理方法添加到路由映射表
func (engine *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	engine.router.addRoute(method, pattern, handler)
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRoute("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
