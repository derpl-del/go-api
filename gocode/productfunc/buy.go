package productfunc

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/derpl-del/go-api.git/gocode/dbadapter"
	"github.com/derpl-del/go-api.git/gocode/emailfunc"
	"github.com/derpl-del/go-api.git/gocode/strcode"
	"github.com/derpl-del/go-api.git/gocode/userfunc"
	"github.com/derpl-del/go-api.git/gocode/utilfunc"
)

var validation bool

//BuyProduct func
func BuyProduct(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var Request strcode.BuyProduct
	json.Unmarshal(ReqBody, &Request)
	userval := userfunc.SelectUserDB("username||'|'||EMAIL||'|'||wallet", Request.UserName)
	usermap := utilfunc.TokenizeWithValue("username|mail|wallet", userval)
	fmt.Println(usermap)
	price := Request.Price * Request.Amount
	if usermap["username"] != "" {
		validation = true
	} else {
		validation = false
	}
	if validation == true {
		isproduct := ProducExistsValidation(Request.ProductName)
		if isproduct == true {
			out, err := ProcessBuyDB(usermap["username"], price, Request.ProductName, Request.Amount)
			if err == nil && out == "success" {
				mailval := strconv.Itoa(Request.Amount) + "|" + strconv.Itoa(Request.Price) + "|" + strconv.Itoa(price)
				fmt.Println(mailval)
				emailfunc.GenerateEmail(usermap["mail"], mailval, "invoice")
				errorcode = "0000"
				errormsg = "success"
			} else {
				errorcode = "0003"
				errormsg = "wallet is not enough"
			}
		} else {
			errorcode = "0002"
			errormsg = "product not found"
		}
	} else {
		errorcode = "0001"
		errormsg = "username not found"
	}
	response := strcode.Response{ErrorCode: errorcode, ErrorMsg: errormsg}
	json.NewEncoder(w).Encode(response)
}

//ProcessBuyDB func
func ProcessBuyDB(username string, price int, productname string, amount int) (string, error) {
	db := dbadapter.OpenConnection()
	defer db.Close()
	//using function
	query := fmt.Sprintf("BEGIN :1 := PROCESSBUY('%v',%v); END;", username, price)
	var out string
	_, err := db.Exec(query, sql.Out{Dest: &out})
	if err != nil {
		return "", err
	}
	//using procedure
	/*
		var strArr string
		db.Exec("begin PROCEDURE1(:1, :2, :3); end;", username, wallet, sql.Out{Dest: &strArr})
		fmt.Printf("%v", strArr)
	*/
	return out, nil
}
