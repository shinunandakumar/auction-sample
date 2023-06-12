// controller packages handles the apis

package controller

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
)

type Controller struct {
	DB *sql.DB
}

type JsonResponse struct {
	Status  int    `json:"status_code"`
	Message string `json:"message"`
	Error   error  `json:"error_code"`
}

func (c *Controller) Healthcheck(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
	// fmt.Fprintf(w, "health is fine\n")
	response := JsonResponse{Status: http.StatusOK, Message: "Health is fine", Error: nil}
	json.NewEncoder(w).Encode(response)

	// fmt.Println("sqlConn", c.DB.Ping())
}

func (c *Controller) GetAds(w http.ResponseWriter, r *http.Request) {
	// can call the c.DB
	statement := "SELECT * FROM adsdata"
	// TODO error handling
	rows, _ := c.DB.Query(statement)

	if rows != nil {
		defer rows.Close()
		for rows.Next() {
			fmt.Fprintf(w, "GetAds is fine\n")
		}
	}
	fmt.Fprintf(w, "nil DAta\n")
	w.WriteHeader(http.StatusOK)
}

func (c *Controller) PostAds(w http.ResponseWriter, r *http.Request) {
	// can call the c.DB
	fmt.Fprintf(w, "PostAds is fine\n")
	// statement := fmt.Sprintf("INSERT INTO adsdata (ad_name,base_price,ad_body,status,crated_date,end_date) VALUES ('%s','%s','%s','%s','%d','%d')", "T V", "12000", "Samsung LED TV 40 INCH", "Active", 1686589438, 1686589438)
	// _, err := c.DB.Exec(statement)
	// if err != nil {
	// 	fmt.Println("aaaaaaa", err)
	// 	fmt.Fprintf(w, "PostAds is not fine\n")
	// }
	// fmt.Fprintf(w, "PostAds is fine\n")
	// fmt.Println("sqlConn", c.DB)
}
