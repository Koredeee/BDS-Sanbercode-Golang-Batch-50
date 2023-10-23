package main

import (
	"fmt"
	"log"
	"net/http"

	"config"

	"github.com/julienschmidt/httprouter"
)

func main() {
	db, e := config.MySQL()

	if e != nil {
		log.Fatal(e)
	}

	eb := db.Ping()
	if eb != nil {
		panic(eb.Error())
	}

	fmt.Println("Succes")

	router := httprouter.New()
	fmt.Println("Server running at port 8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
