package userfunc

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/derpl-del/go-api.git/gocode/endecode"
	"github.com/derpl-del/go-api.git/gocode/readfunc"
	"github.com/derpl-del/go-api.git/gocode/strcode"
	"github.com/derpl-del/go-api.git/gocode/writefunc"
)

//UserVerify func
func UserVerify(w http.ResponseWriter, r *http.Request) {
	req := r.FormValue("req")
	//fmt.Println(req)
	result := endecode.GenerateDe(req)
	duplicate := UserExistsValidation(result)
	if duplicate == true {
		//fmt.Println("ada")
		var jsonvalue strcode.UserInfo
		filename := result + ".json"
		bytevalue := readfunc.ReadFile("gofile/tmpuser/", filename)
		json.Unmarshal(bytevalue, &jsonvalue)
		//fmt.Println(jsonvalue.UserName)
		err := InsUserDB(jsonvalue.UserName, jsonvalue.Wallet, jsonvalue.Mail)
		if err != nil {
			fmt.Println(fmt.Sprintf("%v", err))
		} else {
			writefunc.DeleteFile("gofile/tmpuser/" + filename)
		}
	} else {
		//fmt.Println("tidak")
	}

}
