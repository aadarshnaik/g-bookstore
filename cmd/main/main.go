package main

import "github.com/gorilla/mux"

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/book", getBooks).Methods("GET")
}
