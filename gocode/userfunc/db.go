package userfunc

import (
	"database/sql"
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

//DuplicateDB func
func DuplicateDB(username string, email string) string {
	db := dbadapter.OpenConnection()
	defer db.Close()
	//using function
	query := fmt.Sprintf("BEGIN :1 := DUPLICATEFUNC('%v','%v'); END;", username, email)
	var out string
	_, err := db.Exec(query, sql.Out{Dest: &out})
	if err != nil {
		return fmt.Sprintf("%v", err)
	}
	//using procedure
	/*
		var strArr string
		db.Exec("begin PROCEDURE1(:1, :2, :3); end;", username, wallet, sql.Out{Dest: &strArr})
		fmt.Printf("%v", strArr)
	*/
	return out
}
