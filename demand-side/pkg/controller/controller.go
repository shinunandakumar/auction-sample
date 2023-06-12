// controller packages handles the apis

package controller

import (
	"database/sql"
	"fmt"
	"net/http"
)

type Controller struct {
	DB *sql.DB
}

func (c *Controller) Healthcheck(w http.ResponseWriter, r *http.Request) {
	// can call the c.DB
	fmt.Fprintf(w, "health is fine\n")
	fmt.Println("sqlConn", c.DB.Ping())
}
