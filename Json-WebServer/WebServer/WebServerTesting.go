package main

import (
	f "fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	f.Fprintln(w, "what's up peasants")
}

func home(responseWriter http.ResponseWriter, r *http.Request) {
	f.Fprintln(responseWriter, "You're at home dumbass!")
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		f.Fprintln(w, "halo")
	})

	// function http.HandleFunc() --> untuk routing aplikasi web (penentuan aksi ketika url tertentu diakses oleh user)
	// terdapat 2 parameter (rute yang diinginkan [yang akan diakses oleh user], callback/aksi ketika rute tersebut diakses (func(w http.ResponseWriter, r *http.Request))
	http.HandleFunc("/index", index)

	http.HandleFunc("/home", home)

	f.Println("starting web server at http://localhost:8080/")

	http.ListenAndServe(":8080", nil)
}
