package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vingarcia/simplehttp"
	shttp "github.com/vingarcia/simplehttp"
)

func main() {
	fmt.Println("serving on port 8080...")
	http.ListenAndServe(":8080", simplehttp.Handler{
		Funcs: []shttp.Middleware{

			func(rw shttp.ResponseWriter, r *http.Request, next func()) {
				fmt.Printf("receiving request %s %s\n", r.Method, r.URL.Path)
				next()
				fmt.Printf("finishing request %s %s: %d\n", r.Method, r.URL.Path, rw.Status())
			},

			func(rw shttp.ResponseWriter, r *http.Request, next func()) {
				fmt.Printf("authenticating...\n")
				next()
			},

			func(rw shttp.ResponseWriter, r *http.Request, next func()) {
				fmt.Printf("measuring time spent...\n")
				startTime := time.Now()

				next()

				fmt.Printf("time span: %vs\n", time.Since(startTime).Seconds())
			},

			func(rw shttp.ResponseWriter, r *http.Request, next func()) {
				fmt.Fprintf(rw, "Welcome to my server")
			},
		},
	})
}
