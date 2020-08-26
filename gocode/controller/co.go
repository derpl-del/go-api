package controller

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/go-api.git/gocode/emailfunc"
	"github.com/derpl-del/go-api.git/gocode/productfunc"
	"github.com/derpl-del/go-api.git/gocode/userfunc"
	"github.com/gorilla/mux"
)

var r *mux.Router

//CtrFunc for find
func CtrFunc() {
	fmt.Println("morning")
	r = mux.NewRouter()
	r.HandleFunc("/hello", productfunc.HelloName)
	r.HandleFunc("/api/v1/generatesender", emailfunc.GenerateEmailSender).Methods("POST")
	r.HandleFunc("/api/v1/verifyuser", userfunc.UserVerify)
	r.HandleFunc("/api/v1/createuser", userfunc.CreateUser).Methods("POST")
	r.HandleFunc("/api/v1/buyproduct", productfunc.BuyProduct).Methods("POST")
	r.HandleFunc("/api/v1/createproduct", productfunc.CreateProduct).Methods("POST")
	r.HandleFunc("/api/v1/getallproduct", productfunc.GetViewProduct)
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
