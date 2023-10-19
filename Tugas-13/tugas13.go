package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
)

type NilaiMahasiswa struct {
	Nama        string `json:"nama"`
	MataKuliah  string `json:"mata_kuliah"`
	IndeksNilai string `json:"indeks_nilai"`
	Nilai       int    `json:"nilai"`
	ID          uint   `json:"id"`
}

var nilaiNilaiMahasiswa = []NilaiMahasiswa{}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()
		if !ok {
			w.Write([]byte("Username atau Passowrd tidak boleh kosong"))
			return
		}

		if username == "mahasiswa" && password == "mahasiswa" {
			next.ServeHTTP(w, r)
			return
		} else {
			w.Write([]byte("Invalid Username or Password"))
		}

	})
}

func PostNilai(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var Nilai NilaiMahasiswa

	if r.Method == "POST" {
		if r.Header.Get("Content-Type") == "application/json" {
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&Nilai); err != nil {
				log.Fatal(err)
			}
		} else if r.FormValue("nama") != "" && r.FormValue("mata_kuliah") != "" && r.FormValue("nilai") != "" {
			getID := r.PostFormValue("id")
			id, _ := strconv.Atoi(getID)
			nama := r.PostFormValue("nama")
			mataKuliah := r.PostFormValue("mata_kuliah")
			getNilai := r.PostFormValue("nilai")
			nilai, _ := strconv.Atoi(getNilai)

			var indeksNilai string

			if nilai >= 80 {
				indeksNilai = "A"
			} else if nilai >= 70 && nilai < 80 {
				indeksNilai = "B"
			} else if nilai >= 60 && nilai < 70 {
				indeksNilai = "C"
			} else if nilai >= 50 && nilai < 60 {
				indeksNilai = "D"
			} else if nilai < 50 {
				indeksNilai = "E"
			}

			Nilai = NilaiMahasiswa{
				Nama:        nama,
				MataKuliah:  mataKuliah,
				IndeksNilai: indeksNilai,
				Nilai:       nilai,
				ID:          uint(id),
			}
		} else {
			http.Error(w, "Invalid [Bad Request :(]", http.StatusBadRequest)
		}

		nilaiNilaiMahasiswa = append(nilaiNilaiMahasiswa, Nilai)

		dataNilai, _ := json.Marshal(Nilai)
		w.Write(dataNilai)
		return
	}

	http.Error(w, "NOT FOUND", http.StatusNotFound)
}

func getNilai(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		nilaiNilai := nilaiNilaiMahasiswa
		dataNilai, err := json.Marshal(nilaiNilai)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(dataNilai)

		return
	} else {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	server := &http.Server{
		Addr: ":8080",
	}

	// routing
	http.Handle("/post_nilai", Auth(http.HandlerFunc(PostNilai)))
	http.HandleFunc("/nilai_mahasiswa", getNilai)

	// server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()

}
