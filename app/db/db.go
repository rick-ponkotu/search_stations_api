package db

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

//go: embed schema.sql
//var schema string

//NewDB returns mysql driver based *sql.DB.
func NewDBMysql(path string) (*sql.DB, error) {
	db, err := sql.Open("mysql", path)
	if err != nil {
		return nil, err
	}

	return db, nil
}
