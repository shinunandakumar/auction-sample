// module demand-side consist of the apis of
// demand side of the auction service
// entrypoint of the application

package main

import (
	"database/sql"
	"demand-side/pkg/controller"
	"demand-side/pkg/db"
	"fmt"
	"log"
	"net/http"
)

var dbConn *sql.DB

func init() {
	dbConn = db.Conn()

	db.Automigrate(dbConn)
}

func main() {

	fmt.Println("demand side is running")
	defer dbConn.Close()

	// define controller package
	ctr := controller.Controller{
		DB: dbConn,
	}

	// TODO add more handlers
	http.HandleFunc("/health-check", ctr.Healthcheck)

	// http.HandleFunc("/get-bid", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "get bid is fine\n")
	// })
	// http.HandleFunc("/post-bid", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "post bid is fine\n")
	// })
	// TODO change port to env
	if err := http.ListenAndServe(":8001", nil); err != nil {
		log.Fatal(err)
	}
}
