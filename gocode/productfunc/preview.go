package productfunc

import (
	"encoding/json"
	"net/http"

	"github.com/derpl-del/go-api.git/gocode/readfunc"
	"github.com/derpl-del/go-api.git/gocode/strcode"
)

//ProductList var
var ProductList []strcode.ProductInfo

//GetViewProduct func
func GetViewProduct(w http.ResponseWriter, r *http.Request) {
	ProductData := readfunc.GetAllProduct()
	json.NewEncoder(w).Encode(ProductData)
}
