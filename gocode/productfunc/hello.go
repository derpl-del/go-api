package productfunc

import (
	"fmt"
	"net/http"
)

//HelloName func
func HelloName(w http.ResponseWriter, r *http.Request) {
	input := r.FormValue("name")
	text := "hello " + input
	fmt.Fprintf(w, text)
}
