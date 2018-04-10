package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside HelloServer handler")
	fmt.Fprintf(w, "Hello,"+req.URL.Path[1:])
}
func PostServer(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Inside PostServer handler")
	fmt.Fprintf(w, "Post,"+req.URL.Path[1:])
}
func main() {
	http.HandleFunc("/Hello", HelloServer)
	http.HandleFunc("/Post", PostServer)
	fmt.Println("Server start...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
