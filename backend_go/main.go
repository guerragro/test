package main

import (
	// "fmt"
	"log"
	"net/http"
	// "github.com/gorilla/mux"
)
// рабочая версия
func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		log.Print(r)
		log.Print(r.Body)
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Write([]byte("<h1>Hello World</h1>"))
	})
	log.Print("server starting")
	http.ListenAndServe(":8080", nil)
}


// type Req struct{
// 	Username string 'json:"username"'
// 	Password string 'json:"password"'
// }

// func main(){
// 	router := mux.NewRouter()
// 	router.HandleFunc("/", makeRequest)
// 	log.Print("server start")
// 	log.Fatal(http.ListenAndServe(":8080", router))
// }

// func makeRequest(w http.ResponseWriter, r *http.Request){
// 	// mux.Vars(r)
// 	// log.Print(w)
// 	log.Print(r.Body)
// 	w.Header().Set("Access-Control-Allow-Origin", "*")
// 	w.Write([]byte("hello world"))
// 	// req := Req{}
// 	// json.NewDecoder(r.Body).Decode(&req)
// }