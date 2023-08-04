package gee

import (
	"fmt"
	"log"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

//// GetRouter show router of engine
//func (engine *Engine) GetRouter() map[string]HandlerFunc {
//	return engine.router
//}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := GetKey(r.Method, r.URL.Path)
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{
		router: make(map[string]HandlerFunc),
	}
}

// method: GET, POST, etc
// pattern: like /hello
func (engine *Engine) addRouter(method, pattern string, handler HandlerFunc) {
	key := GetKey(method, pattern)
	log.Printf("Route %4s - %s", method, pattern)
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	engine.addRouter("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRouter("POST", pattern, handler)
}

func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}

func GetKey(method, pattern string) string {
	return method + "-" + pattern
}
