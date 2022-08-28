package main

import (
    "fmt"
    "log"
    "net/http"
)

var name string = "STUDENT_NAME_HERE"

func helloHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path != "/hello" {
        http.Error(w, "404 not found.", http.StatusNotFound)
        return
    }

    if r.Method != "GET" {
        http.Error(w, "Method is not supported.", http.StatusNotFound)
        return
    }


    fmt.Fprintf(w, "Hello, %s!", name)
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
    http.Handle("/", fileServer) 
	
    http.HandleFunc("/hello", helloHandler)
	
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
        log.Fatal(err)
    }
}
