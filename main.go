package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin-go-sdk/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

	})
}

func main() {}
