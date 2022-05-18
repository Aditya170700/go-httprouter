package main

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/julienschmidt/httprouter"
	"github.com/stretchr/testify/assert"
)

func TestRouter(t *testing.T) {
	router := httprouter.New()

	router.GET("/", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprint(w, "Test Hello HttpRouter")
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/", nil)

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Test Hello HttpRouter", string(body))
}

func TestRouterParams(t *testing.T) {
	router := httprouter.New()

	router.GET("/products/:id", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		fmt.Fprintf(w, "Product %s", p.ByName("id"))
	})

	recorder := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/products/1", nil)

	router.ServeHTTP(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	assert.Equal(t, "Product 1", string(body))
}
