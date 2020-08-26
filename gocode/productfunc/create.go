package productfunc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/derpl-del/go-api.git/gocode/strcode"
	"github.com/derpl-del/go-api.git/gocode/writefunc"
)

//ProductInfo var
var ProductInfo strcode.ProductInfo
var errorcode string
var errormsg string

//CreateProduct func
func CreateProduct(w http.ResponseWriter, r *http.Request) {
	ReqBody, _ := ioutil.ReadAll(r.Body)
	var ReqProduct strcode.ProductInfo
	json.Unmarshal(ReqBody, &ReqProduct)
	result := CreateProductHandler(ReqProduct.ProductName)
	if result == false && ReqProduct.ProductAmount > 0 && ReqProduct.ProductPrice > 0 {
		path := "gofile/product/" + ReqProduct.ProductName + ".json"
		err := writefunc.WriteFile(path, ReqBody)
		if err != nil {
			errorcode = "9999"
			errormsg = fmt.Sprintf("%v", err)
		} else {
			errorcode = "0000"
			errormsg = "success"
		}
	} else if ReqProduct.ProductAmount <= 0 && ReqProduct.ProductPrice <= 0 {
		errorcode = "0001"
		errormsg = "invalid amount value and price value"
	} else if ReqProduct.ProductAmount <= 0 {
		errorcode = "0002"
		errormsg = "invalid amount value"
	} else if ReqProduct.ProductPrice <= 0 {
		errorcode = "0002"
		errormsg = "invalid price value"
	} else {
		errorcode = "9999"
		errormsg = "duplicate data"
	}
	response := strcode.Response{ErrorCode: errorcode, ErrorMsg: errormsg}
	json.NewEncoder(w).Encode(response)
}

//CreateProductHandler handler for duplicate data
func CreateProductHandler(productname string) bool {
	for _, data := range ProductList {
		if data.ProductName == productname {
			return true
		}
		fmt.Println(fmt.Sprintf("product name : %v", data.ProductName))
	}
	return false
}
