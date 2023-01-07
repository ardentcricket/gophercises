package main

import (
	"io"
	"net/http"
	"log"
)

func handleHello(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello world - nn \n")
}

func main() {
	http.HandleFunc("/hello", handleHello)
	log.Println("Listening @ localhost:8000/hello")
	http.ListenAndServe(":8000", nil)
}