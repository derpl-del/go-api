package dbadapter

import (
	"database/sql"
	"fmt"

	//gordor import
	_ "github.com/godror/godror"
)

//OpenConnection func
func OpenConnection() *sql.DB {
	db, err := sql.Open("godror", "go_api/welcome1@xe")
	if err != nil {
		fmt.Println(err)
		return db
	}
	return db
}
