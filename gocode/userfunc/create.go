package userfunc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"

	"github.com/derpl-del/go-api.git/gocode/emailfunc"
	"github.com/derpl-del/go-api.git/gocode/strcode"
	"github.com/derpl-del/go-api.git/gocode/writefunc"
)

var errorcode string
var errormsg string
var emailRegex = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

//CreateUser func
func CreateUser(w http.ResponseWriter, r *http.Request) {
	var result string
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var ReqUser strcode.UserInfo
	json.Unmarshal(ReqBody, &ReqUser)
	duplicate := UserExistsValidation(ReqUser.UserName)
	if duplicate == true {
		result = "0001"
	} else {
		result = DuplicateDB(ReqUser.UserName, ReqUser.Mail)
		fmt.Println(result)
	}
	isemail := IsEmailValid(ReqUser.Mail)
	if result == "0000" && isemail == true && ReqUser.Wallet > 0 {
		path := "gofile/tmpuser/" + ReqUser.UserName + ".json"
		err := writefunc.WriteFile(path, ReqBody)
		if err != nil {
			errorcode = "0003"
			errormsg = fmt.Sprintf("%v", err)
		} else {
			emailfunc.GenerateEmail(ReqUser.Mail, ReqUser.UserName, "user verify")
			errorcode = "0000"
			errormsg = "success"
		}
	} else if result == "0001" {
		errorcode = "0001"
		errormsg = "user already registered"
	} else if result == "0002" {
		errorcode = "0002"
		errormsg = "mail already registered"
	} else if isemail == false {
		errorcode = "0004"
		errormsg = "invalid mail value"
	} else {
		errorcode = "0005"
		errormsg = "invalid wallet amount"
	}
	response := strcode.Response{ErrorCode: errorcode, ErrorMsg: errormsg}
	json.NewEncoder(w).Encode(response)
}

//IsEmailValid func
func IsEmailValid(mail string) bool {
	if len(mail) < 3 && len(mail) > 254 {
		return false
	}
	return emailRegex.MatchString(mail)
}

//UserExistsValidation handler for duplicate data
func UserExistsValidation(username string) bool {
	filename := "gofile/tmpuser/" + username + ".json"
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		return false
	}
	return true
}
