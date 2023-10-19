package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
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

func log(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Ini dari middleware Log....\n")
		fmt.Println(r.URL.Path)
		next.ServeHTTP(w, r)
	})
}

func CekLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		if r.URL.Query().Get("token") != "12345" {
			fmt.Fprintf(w, "Token tidak tersedia atau salah\n")
			return
		}
		next.ServeHTTP(w, r)
	})
}

func getMovies(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
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

func PostMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Mov Movie

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Mov); err != nil {
				err.Error()
			}
		} else {
			// parse dari form
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			title := r.PostFormValue("title")
			getYear := r.PostFormValue("year")
			year, _ := strconv.Atoi(getYear)
			Mov = Movie{
				ID:    id,
				Title: title,
				Year:  year,
			}
		}

		dataMovie, _ := json.Marshal(Mov) // to byte
		w.Write(dataMovie)                // cetak di browser
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
}

func main() {
	// konfigurasi server
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	// http.Handle("/", log(http.HandlerFunc(getMovies)))
	http.Handle("/", CekLogin(http.HandlerFunc(getMovies)))
	http.HandleFunc("/movies", getMovies)
	http.HandleFunc("/post_movie", PostMovie)

	// if err := http.ListenAndServe(":8080", nil); err != nil {
	// 	log.Fatal(err)
	// }
	// running server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}
