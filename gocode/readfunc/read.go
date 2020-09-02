package readfunc

import (
	"io/ioutil"
	"os"
)

//ReadFile func
func ReadFile(path string, filename string) []byte {
	file := path + filename
	req, err := os.Open(file)
	if err != nil {
		//fmt.Println(err)
		return nil
	}
	defer req.Close()
	byteValue, _ := ioutil.ReadAll(req)
	return byteValue
}
