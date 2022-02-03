package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/the-bets/capitalism-game/pkg/data"
)

func main() {
	// set up http handler to listen on port 3000
	http.HandleFunc("/", landingPageHandler)

	fmt.Println("server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

	// begin reading cik data to []string and writing to as-of-now empty cik_list.txt
	cik_lines, err := read.ReadFile("https://www.sec.gov/Archives/edgar/cik-lookup-data.txt")
	if err != nil {
		log.Fatalf("Error reading from text file: %s\n", err)
		return
	}

	err = write.WriteFile(cik_lines, "cik_list.txt")
	if err != nil {
		log.Fatalf("Error writing to cik_lines.txt: %s", err)
	}
	return
}

func landingPageHandler(w http.ResponseWriter, r *http.Request) {
	_, err := fmt.Fprintf(w, "Welcome! Come back soon for investment fun!")
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		log.Panic(err)
	}
}
