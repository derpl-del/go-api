package controller

import (
	"fmt"
	"net/http"

	"github.com/derpl-del/go-api.git/gocode/getfunc"
	"github.com/gorilla/mux"
)

var r *mux.Router

//CtrFunc for find
func CtrFunc() {
	fmt.Println("morning")
	r = mux.NewRouter()
	r.HandleFunc("/", getfunc.GetViewProduct)
	fmt.Println("server started at localhost:9000")
	http.ListenAndServe(":9000", r)
}
