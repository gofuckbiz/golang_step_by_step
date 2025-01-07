package main

import (
	"fmt"
	"net/http"
)

var cache = make(map[string]string)

func ValueByKey(w http.ResponseWriter, request *http.Request) {

	key := request.URL.Query().Get("key")

	value, exists := cache[key]

	if !exists {
		http.Error(w, "Key not found ", http.StatusNotFound)

	}

	if _, err := fmt.Fprintf(w, "%s", value); err != nil {
		http.Error(w, "Internal server Error", http.StatusInternalServerError)
	}

}

func SetValue(w http.ResponseWriter, request *http.Request) {

	key := request.URL.Query().Get("key")

	value := request.URL.Query().Get("value")

	cache[key] = value

	fmt.Fprintf(w, "Value set for key: %s", key)

}

func DeleteValue(w http.ResponseWriter, request *http.Request) {

	// нужно удалить значение по ключу

	key := request.URL.Query().Get("key")

	if _, exists := cache[key]; exists {
		delete(cache, key)
		fmt.Fprintf(w, "value deleted for key: %s", key)

	} else {
		http.Error(w, "key not found", http.StatusNotFound)
	}

}

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
	http.HandleFunc("/value_key", ValueByKey)
	http.HandleFunc("/set_value", SetValue)
	http.HandleFunc("/delete_key", DeleteValue)

	if err := http.ListenAndServe(":8080", nil); err != nil {

		fmt.Printf("Error starting server:%v\n", err)
		return
	}

}
