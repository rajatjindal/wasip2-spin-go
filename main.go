package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin-go-sdk/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("x-rj-is-here", "y")
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("hello from wasip2 golang"))
	})
}

func main() {}
