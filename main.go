package main

import (
	"encoding/json"
	"net/http"
	"time"
)

type Message struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func handlerHello(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {

		ts := time.Now().Format("20060102150405")

		responseObject := Message{
			Message: "Hello World, current timestamp is: " + ts,
			Code:    200,
		}

		response, err := json.Marshal(responseObject)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.Write(response)
	}

}

func main() {
	http.HandleFunc("/hello", handlerHello)
	http.ListenAndServe(":9090", nil)
}
