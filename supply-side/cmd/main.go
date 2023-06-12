// module supply-side consist of the apis of
// supply side of the auction service
// entrypoint of the application

package main

import (
	"fmt"
	"net/http"
	"supply-side/pkg/db"
)

func init() {
	dbConn := db.Conn()
	defer dbConn.Close()

	db.Automigrate(dbConn)
}

func main() {
	fmt.Println("supply side is called")

	http.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "health is fine\n")
	})
	// TODO add more handlers
	// TODO change port to env
	if err := http.ListenAndServe(":8002", nil); err != nil {
		panic("unable to connect")
	}
}
