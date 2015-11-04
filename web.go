package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", HelloWorldHandler)
	log.Fatal(http.ListenAndServe("127.0.0.1:61389", nil))
}

func HelloWorldHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, from the web!"))
}
