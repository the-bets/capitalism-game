package main

import (
	"fmt"
	"log"
	"net/http"

<<<<<<< HEAD
	"github.com/the-bets/capitalism-game/pkg/data"
=======
	read "github.com/the-bets/capitalism-game/pkg/data_handling/read"
	write "github.com/the-bets/capitalism-game/pkg/data_handling/write"
>>>>>>> 699d1598dd0c2e5ae4a9ce199a334e6644195872
)

func main() {
	// set up http handler to listen on port 3000
	http.HandleFunc("/", landingPageHandler)

	fmt.Println("server listening on port 3000")
	log.Panic(
		http.ListenAndServe(":3000", nil),
	)

	// begin reading cik data to []string and writing to as-of-now empty cik_list.txt
<<<<<<< HEAD
	cik_lines, err := data.ReadFile("https://www.sec.gov/Archives/edgar/cik-lookup-data.txt")
=======
	cik_lines, err := read.ReadFile("https://www.sec.gov/Archives/edgar/cik-lookup-data.txt")
>>>>>>> 699d1598dd0c2e5ae4a9ce199a334e6644195872
	if err != nil {
		log.Fatalf("Error reading from text file: %s\n", err)
		return
	}
<<<<<<< HEAD
	err = data.WriteFile(cik_lines, "cik_list.txt")
=======
	err = write.WriteFile(cik_lines, "cik_list.txt")
>>>>>>> 699d1598dd0c2e5ae4a9ce199a334e6644195872
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
