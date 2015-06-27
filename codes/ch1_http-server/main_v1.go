package main

import (
	"net/http"
)

func main() {

	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		w.Write([]byte("Hello World! Go is awesome"))
	})

	http.ListenAndServe(":7000", nil)

}