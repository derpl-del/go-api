package productfunc

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/derpl-del/go-api.git/gocode/dbadapter"
	"github.com/derpl-del/go-api.git/gocode/emailfunc"
	"github.com/derpl-del/go-api.git/gocode/strcode"
	"github.com/derpl-del/go-api.git/gocode/userfunc"
)

//BuyProduct func
func BuyProduct(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	emailfunc.GenerateEmail("gilsptr@gmail.com", "buy")
	var Request strcode.BuyProduct
	json.Unmarshal(ReqBody, &Request)
	validation := userfunc.ValidationUser(Request.UserName)
	if validation == false {
		isproduct := ProducExistsValidation(Request.ProductName)
		if isproduct == true {
			out, err := ProcessBuyDB(Request.UserName, Request.Wallet, Request.ProductName, Request.Amount)
			if err == nil && out == "success" {
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
func ProcessBuyDB(username string, wallet int, productname string, amount string) (string, error) {
	db := dbadapter.OpenConnection()
	defer db.Close()
	//using function
	query := fmt.Sprintf("BEGIN :1 := PROCESSBUY('%v',%v); END;", username, wallet)
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
