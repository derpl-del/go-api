package userfunc

import (
	"fmt"

	"github.com/derpl-del/go-api.git/gocode/dbadapter"
)

//ValidationUser func
func ValidationUser(username string) bool {
	db := dbadapter.OpenConnection()
	defer db.Close()
	QueryString := fmt.Sprintf("select username from USER_INFO where USERNAME = '%v'", username)
	rows, err := db.Query(QueryString)
	if err != nil {
		return false
	}
	defer rows.Close()
	var Rs string
	for rows.Next() {
		rows.Scan(&Rs)
	}
	if Rs != "" {
		return false
	}
	return true
}

//SelectUserDB func
func SelectUserDB(component string, username string) string {
	db := dbadapter.OpenConnection()
	defer db.Close()
	QueryString := fmt.Sprintf("select %v from USER_INFO where USERNAME = '%v'", component, username)
	//fmt.Println(QueryString)
	rows, err := db.Query(QueryString)
	if err != nil {
		//fmt.Println(fmt.Sprintf("%v", err))
		return ""
	}
	defer rows.Close()
	var Rs string
	for rows.Next() {
		rows.Scan(&Rs)
	}
	return Rs
}

//InsUserDB func
func InsUserDB(username string, wallet int, email string) error {
	db := dbadapter.OpenConnection()
	defer db.Close()
	QueryString := fmt.Sprintf("INSERT INTO USER_INFO ( USERNAME, WALLET,EMAIL ,CREATED_DATE, LAST_UPDATE) VALUES ( '%v','%v' ,'%v', SYSDATE,SYSDATE)", username, wallet, email)
	rows, err := db.Query(QueryString)
	if err != nil {
		return err
	}
	defer rows.Close()
	return nil
}
