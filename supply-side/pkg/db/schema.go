// db package consist of the db schema and functions

package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type AdsData struct {
	Id          int
	AdName      string
	BasePrice   string
	AdBody      string
	Status      string
	CreatedDate int
	EndDate     int
}

func Conn() *sql.DB {
	// TODO change this to env
	db, err := sql.Open("mysql", "auctiontest:auctiontest@tcp(db:3306)/auctiondb")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

func Automigrate(db *sql.DB) {

	var AdsCreateTable = `CREATE TABLE IF NOT EXISTS adsdata
							(
								id int,
								ad_name CHAR(255),
								base_price CHAR(255),
								ad_body text,
								status CHAR(255),
								crated_date int,
								end_date int
							)`

	err := db.QueryRow(AdsCreateTable)
	fmt.Println("err----------", err)
}
