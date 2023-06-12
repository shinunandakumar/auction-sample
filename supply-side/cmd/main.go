// module supply-side consist of the apis of
// supply side of the auction service
// entrypoint of the application

package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"supply-side/pkg/controller"
	"supply-side/pkg/db"
)

var dbConn *sql.DB

func init() {
	dbConn = db.Conn()

	db.Automigrate(dbConn)
}

func main() {
	fmt.Println("supply side is called")
	defer dbConn.Close()

	ctr := controller.Controller{
		DB: dbConn,
	}

	http.HandleFunc("/health-check", ctr.Healthcheck)
	http.HandleFunc("/get-ads", ctr.GetAds)
	http.HandleFunc("/post-ads", ctr.PostAds)

	// http.HandleFunc("/health-check", func(w http.ResponseWriter, r *http.Request) {
	// 	fmt.Fprintf(w, "health is fine\n")
	// })
	// TODO add more handlers
	// TODO change port to env
	if err := http.ListenAndServe(":8002", nil); err != nil {
		panic("unable to connect")
	}
}
