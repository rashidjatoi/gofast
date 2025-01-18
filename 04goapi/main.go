package main

import (
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Building Rest APIS")

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "Hello World")
	})

	if err := http.ListenAndServe("localhost:8080", mux); err != nil {
		fmt.Print(err.Error())
	}
}
