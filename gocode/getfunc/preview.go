package getfunc

import (
	"fmt"
	"net/http"
)

//GetViewProduct func
func GetViewProduct(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	fmt.Println("Endpoint Hit: homePage")
}
