package helper

type slice struct {

}

var Slice slice

func (c *slice) Search(key interface{}, haystack []interface{})bool{
	var result = false
	if haystack != nil {
		for _, v := range haystack {
			if v == key {
				result = true
				break
			}
		}
	}

	return result
}
