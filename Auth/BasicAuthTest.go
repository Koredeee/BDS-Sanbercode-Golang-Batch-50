package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Movie struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Year  int    `json:"year"`
}

func Movies() []Movie {
	movs := []Movie{
		{1, "Thor", 2012},
		{2, "Pembunuhan Munir", 2000},
		{3, "G30SPKI", 1980},
	}
	return movs
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Password tidak boleh kosong"))
			return
		}

		if username == "admin" && password == "admin" {
			next.ServeHTTP(w, r)
			return
		}

		w.Write([]byte("Username atau password tidak sesuai"))
		return
	})
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		movies := Movies()
		dataMovies, err := json.Marshal(movies)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		// w.Header().Set("Content-Type", "application/json")
		// w.WriteHeader(http.StatusOK)
		// w.Write(dataMovies)
		// return

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataMovies)
		// w.Write([]byte("<h1>Anda berhasil mengakses fungsi getMovies() </h1>"))
		return
	}

	http.Error(w, "ERROR....", http.StatusNotFound)
}

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/", Auth(http.HandlerFunc(getMovies)))

	// server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}
