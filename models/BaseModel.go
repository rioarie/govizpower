package models

import (
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"log"
)

type BaseModel struct{}

// Query method to MySQL Database
// @param  string statement
// @return (*sql.Rows, error)
func (b *BaseModel) Query(statement string) (*sql.Rows, error) {
	db, err := sql.Open("mysql", "febrian:myb11gd4t4@tcp(150.107.149.81:6033)/simbad_data")

	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Open doesn't open a connection. Validate DSN data:
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Execute the query
	rows, err := db.Query(statement)
	if err != nil {
		return nil, err
	} else {
		return rows, nil
	}
}
