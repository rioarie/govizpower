package models

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type BaseModel struct{}

// Query method to MySQL Database
// @param  string statement
// @return (*sql.Rows, error)
func (b *BaseModel) Query(statement string) (*sql.Rows, error) {
	db, err := sql.Open("mysql", "root@/simbad_data")

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
