package helper

type _string struct {

}

var String _string


func (s *_string) DeleteStart(str string, del string)(string,bool){
	var adjustment = 0
	var ok = false
	var result = ""
	if len(str) < len(del){
		ok = false
	}else{
		strB := []byte(str)
		for i, d := range del {
			for k, st := range str{
				if i==k {
					if d==st {
						strB = append(strB[:i-adjustment], strB[i-adjustment+1:]...)
						adjustment++
						ok = true

					}else{
						return "", false
					}
				}
			}
		}
		result = string(strB)
	}

	return result, ok
}

func (s *_string) DeleteEnd(str string, del string)(string,bool){
	var adjustment = 0
	var ok = false
	var result = ""
	if sls := len(str)-len(del);len(str) < len(del){
		ok = false
	}else{
		strB := []byte(str)

		for k, st := range str {
			for i, d := range del {

				if i+sls == k {
					if d==st {
						strB = append(strB[:k-adjustment], strB[k-adjustment+1:]...)
						adjustment++
						ok = true
					}else{
						return "", false
					}
				}
			}
		}
		result = string(strB)
	}
	return result, ok
}
