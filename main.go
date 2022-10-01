package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		fmt.Fprintf(w, "Welcome to the Go server, do share your Twitter in the readme lol 😄")
	})
	
	http.ListenAndServe(":3000", nil)
}
