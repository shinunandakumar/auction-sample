// db package consist of the db schema and functions

package db

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

type BidData struct {
	Id           int
	AdId         int
	BidderName   string
	BiddingPrice string
	Status       string
	CreatedDate  int
	EndDate      int
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

	var BidCreateTable = `CREATE TABLE IF NOT EXISTS biddata
						(
							id int,
							ad_id int,
							bidder_name CHAR(255),
							bidder_price text,
							status CHAR(255),
							crated_date int,
							end_date int
						)`

	_ = db.QueryRow(BidCreateTable)

	// TODO error handling

}
