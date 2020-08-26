package writefunc

import (
	"io/ioutil"
	"os"
)

//WriteFile func
func WriteFile(path string, input []byte) error {
	err := ioutil.WriteFile(path, input, 0777)
	if err != nil {
		return err
	}
	return nil
}

//DeleteFile func
func DeleteFile(path string) {
	os.Remove(path)
}
