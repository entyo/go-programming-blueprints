package main

import (
	"encoding/json"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/entyo/go-programming-blueprints/meander"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())

	meander.APIKey = os.Getenv("GOOGLE_PLACES_API_KEY")

	http.HandleFunc("/journeys", withCORS(func(w http.ResponseWriter, r *http.Request) {
		respond(w, r, meander.Journeys)
	}))

	http.HandleFunc("/recommendations", withCORS(func(w http.ResponseWriter, r *http.Request) {
		q := &meander.Query{
			Journey: strings.Split(r.URL.Query().Get("journey"), "|"),
		}
		q.Lat, _ = strconv.ParseFloat(r.URL.Query().Get("lat"), 64)
		q.Lng, _ = strconv.ParseFloat(r.URL.Query().Get("lng"), 64)
		q.Radius, _ = strconv.Atoi(r.URL.Query().Get("radius"))
		q.CostRangeStr = r.URL.Query().Get("cost")
		places := q.Run()
		respond(w, r, places)
	}))
	http.ListenAndServe(":8080", http.DefaultServeMux)
}

// respond は任意のデータをjsonエンコードし、http.RequestWriterに書き出す
func respond(w http.ResponseWriter, r *http.Request, data []interface{}) error {
	// publicなデータを格納するための、元データと同じ長さのスライス
	publicData := make([]interface{}, len(data))
	for i, d := range data {
		publicData[i] = meander.Public(d)
	}
	return json.NewEncoder(w).Encode(publicData)
}

// cors対応
func withCORS(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		f(w, r)
	}
}
