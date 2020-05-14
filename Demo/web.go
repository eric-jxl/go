package main

import (
	"fmt"
	"log"
	"net/http"
)

func hander(Writer http.ResponseWriter,request *http.Request)  {
	fmt.Fprintln(Writer,"<h1>Hello World!</h1>")
	fmt.Fprintf(Writer, "<h1>Hi there, I love %s!</h1>", request.URL.Path[1:])
	log.Print(request.URL.Path[1:])
}

func main()  {
	http.HandleFunc("/",hander)
	http.ListenAndServe(":9000",nil)
}
