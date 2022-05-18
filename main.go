package main

import (
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func main() {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Hello HttpRouter")
	})
	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Product %s", p.ByName("id"))
	})

	server := http.Server{
		Handler: router,
		Addr:    "localhost:8080",
	}

	server.ListenAndServe()
}
