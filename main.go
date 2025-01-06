package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, request *http.Request) {

	fmt.Fprintf(w, "Hello go")
}

func headers(w http.ResponseWriter, request *http.Request) {

	for name, headers := range request.Header {
		for _, h := range headers {
			fmt.Fprintf(w, "%v: %v\n", name, h)
		}
	}
}

func main() {

	http.HandleFunc("/hello", hello)
	http.HandleFunc("/headers", headers)

	http.ListenAndServe(":8080", nil)

}
