package main

import (
	"fmt"
	"log"
	"net/http"
)

// func handle(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, "Привет")
// }

// func main() {
// 	http.HandleFunc("/", handle)
// 	log.Print("server starting")
// 	http.ListenAndServe(":8080", nil)
// }

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// r.Header.Add("Access-Control-Allow-Origin", "*")
		r.Header.Get("Access-Control-Allow-Origin : *")
		fmt.Fprint(w, "Привет")
		log.Print(r)

	})
	log.Print("server starting")
	http.ListenAndServe(":8080", nil)
}