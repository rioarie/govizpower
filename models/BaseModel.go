package models

import (
	"database/sql"
	"govizpower/define"
	_ "github.com/go-sql-driver/mysql"
)

type BaseModel struct{}

// Query method to MySQL Database
// @param  string statement
// @return (*sql.Rows, error)
func (b *BaseModel) Query(statement string) (*sql.Rows, error) {
    
    def := define.Init()
    
    db, err := sql.Open("mysql", def.Dburl)
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
