package main

import (
	"net/http"

	spinhttp "github.com/fermyon/spin-go-sdk/http"
	"github.com/fermyon/spin-go-sdk/variables"
)

// const prompt = `<<SYS>>
// You are a bot that generates sentiment analysis responses. Respond with a single positive, negative, or neutral.
// <</SYS>>
// [INST]
// Follow the pattern of the following examples:

// User: Hi, my name is Bob
// Bot: neutral

// User: I am so happy today
// Bot: positive

// User: I am so sad today
// Bot: negative
// [/INST]

// User: {SENTENCE}`

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		// llm infer test
		// result, err := llm.Infer(llm.Llama2Chat, nprompt, nil)
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// kv test
		// store, err := kv.OpenDefault()
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// err = store.Set("testme", []byte("oksure"))
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }

		// outgoing test
		// resp, err := spinhttp.Get("https://google.com")
		// if err != nil {
		// 	http.Error(w, err.Error(), http.StatusInternalServerError)
		// 	return
		// }
		// defer resp.Body.Close()

		// _, err = io.ReadAll(resp.Body)
		// if err != nil {
		// 	http.Error(w, "err.Error()", http.StatusInternalServerError)
		// 	return
		// }

		_, err := variables.Get("password")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "text/plain")
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("hello"))
	})
}

func main() {}
