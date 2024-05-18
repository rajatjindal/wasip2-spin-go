package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin-go-sdk/http"
	"github.com/fermyon/spin-go-sdk/wit"
)

var _ = wit.Wit

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {

	})
}

func main() {}
