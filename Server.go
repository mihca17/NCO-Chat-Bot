package main

import (
	"log"
	"net/http"
)

func Server(address string, port string) error {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println("Server started on", address+":"+port)
		http.FileServer(http.Dir(".")).ServeHTTP(w, r)
	})
	err := http.ListenAndServe(address+":"+port, nil)
	if err != nil {
		log.Fatalln("There was a problem during starting server: ", err)
		return err
	}
	return nil
}
