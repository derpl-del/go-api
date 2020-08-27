package utilfunc

import (
	"strings"
)

//TokenizeWithValue func
func TokenizeWithValue(parameter string, value string) map[string]string {
	svalue := strings.Split(value, "|")
	sparameter := strings.Split(parameter, "|")
	result := make(map[string]string)
	for i := 0; i <= (len(sparameter) - 1); i++ {
		result[sparameter[i]] = svalue[i]
	}
	return result
}
