package main

import (
	"io"
	"net/http"

	spinhttp "github.com/fermyon/spin-go-sdk/http"
)

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		resp, err := spinhttp.Get("https://google.com")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		raw, err := io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "err.Error()", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(resp.StatusCode)
		w.Write(raw)
	})
}

func main() {}
