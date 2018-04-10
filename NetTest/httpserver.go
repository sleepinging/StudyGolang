package main

import (
	"fmt"
	"log"
	"net/http"
)

func HelloServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello")
}
func PostServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Post")
}
func DefaultServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Default")
}
func main() {
	http.HandleFunc("/Hello", HelloServer)
	http.HandleFunc("/Post", PostServer)
	http.HandleFunc("/", DefaultServer)
	fmt.Println("Server start...")
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal("ListenAndServe: ", err.Error())
	}
}
