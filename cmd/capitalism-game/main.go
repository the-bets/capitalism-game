package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", landingPageHandler)

	fmt.Println("server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "hello world")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
