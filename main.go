package main

import (
	"log"
	"net/http"
)

func main() {

	app := NewRouter()
	tmp := NewTemplatesHandler()

	app.mux.Handle("GET /", http.HandlerFunc(tmp.IndexView))

	log.Fatal(app.ListenAndServe(":8080"))
}
