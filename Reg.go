package Helper

import (
	"fmt"
	"regexp"
)

type reg struct{

}

var Reg reg

func (r *reg)Find(reg string, text string)[]string{
	var result = make([]string, 0)
	re := regexp.MustCompile(reg)
	str := fmt.Sprintf("%s", re.Find([]byte(text)))
	if len(str)>0 {
		result = append(result, str)
	}

	return result
}

func (r *reg)FindAll(reg string, text string)[]string{
	var result = make([]string, 0)
	re := regexp.MustCompile(reg)

	for _, r := range re.FindAll([]byte(text), -1){
		result = append(result, string(r))
	}

	return result
}
