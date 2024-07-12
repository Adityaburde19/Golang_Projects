package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(resp http.ResponseWriter, req *http.Request) {
	if err := req.ParseForm(); err != nil {
		fmt.Fprintf(resp, "ParseForm() err: %v", err)
	}
	fmt.Fprintf(resp, "POST request successful")
	name := req.FormValue("name")
	address := req.FormValue("address")
	fmt.Fprintf(resp, "Name = %s\n", name)
	fmt.Fprintf(resp, "Address = %s\n", address)
}

func helloHandler(resp http.ResponseWriter, req *http.Request) {
	if req.URL.Path != "hello" {
		http.Error(resp, "404 not found", http.StatusNotFound)
		return
	}

	if req.Method != "GET" {
		http.Error(resp, "method is not supported", http.StatusNotFound)
	}

	fmt.Println(resp, "Hello")
}

func main() {
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Starting server at port 8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
