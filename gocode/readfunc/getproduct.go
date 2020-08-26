package readfunc

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/derpl-del/go-api.git/gocode/strcode"
)

//GetAllProduct func
func GetAllProduct() strcode.ProductList {
	ProductList := []strcode.ProductInfo{}
	File, err := ioutil.ReadDir("gofile/product/")
	if err != nil {
		fmt.Println(err)
	}
	for _, files := range File {
		bytevalue := ReadFile("gofile/product/", files.Name())
		var jsonvalue strcode.ProductInfo
		json.Unmarshal(bytevalue, &jsonvalue)
		ProductList = append(ProductList, jsonvalue)
	}
	response := strcode.ProductList{}
	response.ProductList = ProductList
	return response
}
