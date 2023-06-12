// module demand-side consist of the apis of
// demand side of the auction service
// entrypoint of the application

package main

import (
	"demand-side/pkg/db"
	"fmt"
	"log"
	"net/http"
)

func init() {
	dbConn := db.Conn()
	defer dbConn.Close()

	db.Automigrate(dbConn)
}

func main() {

	fmt.Println("demand side is running")

	// TODO add more handlers
	http.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "health is fine\n")
	})
	http.HandleFunc("/get-bid", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "get bid is fine\n")
	})
	http.HandleFunc("/post-bid", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "post bid is fine\n")
	})
	// TODO change port to env
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}
