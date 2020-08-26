package productfunc

import "github.com/derpl-del/go-api.git/gocode/readfunc"

//ProducExistsValidation func
func ProducExistsValidation(productname string) bool {
	filename := productname + ".json"
	err := readfunc.ProductHandler("gofile/product/", filename)
	if err != nil {
		return false
	}
	return true
}
