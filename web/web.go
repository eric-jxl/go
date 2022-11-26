package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "<h1>Hello World!</h1>")
	fmt.Fprintf(w, "<h2>Hi there, I love %s!</h2>", r.URL.Path[1:])
	log.Print(r.URL.Path[1:])
	log.Print(r.Header.Get("User-Agent"))
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":9000", nil)
}
