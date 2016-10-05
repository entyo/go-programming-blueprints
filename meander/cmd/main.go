package main

import (
	"encoding/json"
	"net/http"
	"runtime"

	"github.com/entyo/go-programming-blueprints/meander"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	// TODO:
	//   meander.APIKey = "***"

	http.HandleFunc("/journeys", func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	})

	http.ListenAndServe(":8080", http.DefaultServeMux)
}

// respond は任意のデータをjsonエンコードし、http.RequestWriterに書き出す
func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	return json.NewEncoder(w).Encode(data)
}
