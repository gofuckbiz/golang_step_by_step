package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, request *http.Request) {

	if _, err := fmt.Fprintf(w, "Hello go"); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

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

	if err := http.ListenAndServe(":8080", nil); err != nil {

		fmt.Printf("Error starting server:%v\n", err)
		return
	}

}
