package main

import (
	"io"
	"net/http"
	"strings"

	spinhttp "github.com/fermyon/spin-go-sdk/http"
	"github.com/fermyon/spin-go-sdk/kv"
	"github.com/fermyon/spin-go-sdk/llm"
)

const prompt = `<<SYS>>
You are a bot that generates sentiment analysis responses. Respond with a single positive, negative, or neutral.
<</SYS>>
[INST]
Follow the pattern of the following examples:

User: Hi, my name is Bob
Bot: neutral

User: I am so happy today
Bot: positive

User: I am so sad today
Bot: negative
[/INST]

User: {SENTENCE}`

func init() {
	spinhttp.Handle(func(w http.ResponseWriter, r *http.Request) {
		nprompt := strings.ReplaceAll(prompt, "{SENTENCE}", r.URL.Query().Get("sentence"))
		result, err := llm.Infer(llm.Llama2Chat, nprompt, nil)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		store, err := kv.OpenDefault()
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		err = store.Set("testme", []byte("oksure"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		resp, err := spinhttp.Get("https://google.com")
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		defer resp.Body.Close()

		_, err = io.ReadAll(resp.Body)
		if err != nil {
			http.Error(w, "err.Error()", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(resp.StatusCode)
		w.Write([]byte(result.Text))
	})
}

func main() {}
