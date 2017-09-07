package main

import (
	"fmt"
	"io"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello world!")
}

func main() {
	http.HandleFunc("/", hello)
	fmt.Println("server running http://0.0.0.0:8000")
	http.ListenAndServe(":8000", nil)
}
