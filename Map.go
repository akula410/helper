package helper

type _map struct {

}

var Map _map


func (m *_map) Element(key string, haystack map[interface{}]interface{}, dft interface{}) interface{}{
	if result, ok := haystack[key]; ok {
		return result
	}

	return dft
}

func (m *_map) OfElement(name string, value string, haystack []map[interface{}]interface{}, dft map[interface{}]interface{}) map[interface{}]interface{}{
	if haystack != nil {
		for _, item := range haystack {
			for itemName, itemValue := range item {
				if itemName == name && itemValue == value {
					dft = item
					break
				}
			}
		}
	}

	return dft
}
