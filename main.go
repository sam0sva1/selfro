package main

import (
	"net/http"
	"selfro/router"
)

func main() {
	r := router.Start()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello World"))
	})
	r.Get("/valera/da", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("VALERA DA"))
	})
	r.Get("/valera", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("VALERA"))
	})
	r.Get("/nataha", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("NATAHA"))
	})

	r.Run()
}
