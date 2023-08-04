package main

import (
	"fmt"
	"log"
	"net/http"

	"gee"
)

func main() {
	r := gee.New()

	r.GET("/", indexHandler)
	r.GET("/hello", helloHandler)

	log.Fatal(r.Run(":8400"))
}

// handler echoes the Path component of the request URL r.URL.Path
func indexHandler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}

// handler echoes r.URL.Header
func helloHandler(w http.ResponseWriter, r *http.Request) {
	for k, v := range r.Header {
		_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
	}
}
