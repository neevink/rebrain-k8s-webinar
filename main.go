package main

import (
	"fmt"
	"net/http"
	"os"
)

var count uint64

func main() {
	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		count++
		fmt.Fprintf(w, "New hi from %s #%d\n", hostname, count)
	})

	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, "OK")
	})

	fmt.Print("Starting listening on 8080 port")
	http.ListenAndServe(":8080", nil)
}
