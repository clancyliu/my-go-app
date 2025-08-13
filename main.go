package main

import (
	"fmt"
	"net/http"
)

func main() {
	serveMux := http.NewServeMux()

	serveMux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println(request.Method)
		fmt.Fprintf(writer, "Hello World")
	})

	http.ListenAndServe(":8080", serveMux)
}
