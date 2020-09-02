package readfunc

import (
	"os"
)

//ProductHandler func
func ProductHandler(path string, filename string) error {
	file := path + filename
	_, err := os.Open(file)
	if err != nil {
		//fmt.Println(err)
		return err
	}
	return nil
}
