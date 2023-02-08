package main

import (
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	mux.HandleFunc("/slot1", slot1)
	mux.HandleFunc("/slot2", slot2)
	log.Println("ПРОИСХОДИТ ЗАПУСК")
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	w.Write([]byte("КАЗИНО ДЛЯ НЕГРОВ"))
}

func slot1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("СЛОТ ПЕРВЫЙ"))
}

func slot2(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("СЛОТ ВТОРОЙ ОСТОРОЖНО НЕГРЫ!!"))
}
