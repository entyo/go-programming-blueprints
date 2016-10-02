package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var port = flag.String("port", "8081", "Webサイトのアドレス")
	flag.Parse()

	mux := http.NewServeMux()
	handler := http.StripPrefix("/", http.FileServer(http.Dir("public")))
	mux.Handle("/", handler)

	log.Println("Webサイトのポート", ":"+*port)
	http.ListenAndServe(":"+*port, mux)
}
