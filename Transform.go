package Helper

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

type transform struct {

}

var Transform transform

func (t *transform) ToString(value interface{}, def string)string{
	if value != nil {
		switch reflect.TypeOf(value).Kind() {
		case reflect.String:
			def = reflect.ValueOf(value).String()
		case reflect.Slice:
			result := reflect.ValueOf(value)
			for i := 0; i < result.Len(); {
				def = fmt.Sprintf("%v", result.Index(i).Interface())
				break
			}
		}
	}
	return def
}

func (t *transform) ToInt(value interface{}, def int64)int64{
	var err error
	var ok bool
	var prf int

	if value != nil {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Int:
			def = value.(int64)
		case reflect.Slice:
			result := reflect.ValueOf(value)
			for i := 0; i < result.Len(); {
				def = t.ToInt(result.Index(i).Interface(), def)
				break
			}
		case reflect.String:
			result := fmt.Sprintf("%v", value)
			if ok, err = regexp.Match(`^\d+$`, []byte(result)); ok && err == nil {
				prf, err = strconv.Atoi(result)
				if err == nil {
					def = int64(prf)
				}else{
					panic(err)
				}
			}else if err != nil{
				panic(err)
			}
		case reflect.Struct:
			def = t.ToInt(fmt.Sprintf("%v", value), def)
		}

	}
	return def
}

func (t *transform) ToFloat(value interface{}, def float64)float64{
	var err error
	var ok bool

	if value != nil {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Float64:
			def = reflect.ValueOf(value).Float()
		case reflect.Slice:
			result := reflect.ValueOf(value)
			for i := 0; i < result.Len(); {
				def = t.ToFloat(result.Index(i).Interface(), def)
				break
			}
		case reflect.String:
			result := fmt.Sprintf("%v", value)
			if ok, err = regexp.Match(`^\d+\.?\d*$`, []byte(result)); ok && err == nil {
				def, err = strconv.ParseFloat(result, 64)
				if err != nil {
					panic(err)
				}
			}else if err != nil{
				panic(err)
			}
		case reflect.Struct:
			def = t.ToFloat(fmt.Sprintf("%v", value), def)
		}
	}
	return def
}

func (t *transform) ToSlice(value interface{}, def []interface{})[]interface{}{
	if value != nil {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Slice:
			result := reflect.ValueOf(value)
			def = make([]interface{}, 0)

			for i := 0; i < result.Len(); i++ {
				def = append(def, result.Index(i).Interface())
			}
		}
	}

	return def
}

func (t *transform) ToMap(value interface{}, def map[string]interface{})map[string]interface{}{
	if value != nil {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Map:
			def = value.(map[string]interface{})
		case reflect.Slice:
			result := reflect.ValueOf(value)
			for i := 0; i < result.Len(); {
				def = t.ToMap(result.Index(i).Interface(), def)
				break
			}
		}
	}

	return def
}

func (t *transform) ToBool(value interface{}, def bool)bool{
	if value != nil {
		switch reflect.TypeOf(value).Kind() {
		case reflect.Bool:
			def = value.(bool)
		case reflect.Slice:
			result := reflect.ValueOf(value)
			for i := 0; i < result.Len();{
				def = t.ToBool(result.Index(i), def)
				break
			}
		case reflect.String:
			result := strings.ToLower(fmt.Sprintf("%v", value))
			if result == "1" || result == "true" {
				def = true
			}else{
				def = false
			}

		case reflect.Struct:
			def = t.ToBool(fmt.Sprintf("%v", value), def)
		}

	}
	return def
}

func (t *transform) JsonToInterface(data []byte) interface{}{
	var jsonMap map[string]*json.RawMessage
	var jsonSlice []*json.RawMessage
	var jsonScalar *json.RawMessage

	if json.Unmarshal(data, &jsonMap) != nil {
		if json.Unmarshal(data, &jsonSlice) != nil {
			if json.Unmarshal(data, &jsonScalar) != nil {
				return ""
			}
		}
	}


	if jsonMap != nil {
		var resultMap = make(map[string]interface{})
		for i, r := range jsonMap{
			resultMap[i] = t.JsonToInterface(*r)
		}
		return resultMap
	}

	if jsonSlice != nil && len(jsonSlice)>0 {
		var resultSlice = make([]interface{}, 0)
		for _, r := range jsonSlice {
			resultSlice = append(resultSlice, t.JsonToInterface(*r))
		}
		return resultSlice
	}

	return t.ByteToInterface(*jsonScalar)
}

func (t *transform)ByteToInterface(data []byte)interface{}{
	var dataTmp = fmt.Sprintf("%s", data)
	if ok, _ := regexp.Match(`^\"`, data); ok {
		var result string
		result, ok = String.DeleteStart(dataTmp, `"`)
		if ok {
			result, ok = String.DeleteEnd(result, `"`)
		}

		return result
	}

	if dataTmp == "true" {
		return true
	}

	if dataTmp == "false" {
		return false
	}

	if ok, _ := regexp.Match(`^\d+\.?\d*$`, data); ok {
		return dataTmp
	}

	return nil
}


func (t *transform)MapToJson(data interface{})string{
	var result string
	if data != nil {
		var JSON []byte
		var err error
		JSON, err = json.Marshal(data)
		if err != nil {
			panic(err)
		}
		result = string(JSON)
	}
	return result
}

func (t *transform)MapToGetParams(data interface{})string{
	req, err := http.NewRequest("GET", "", nil)
	if err != nil {
		panic(err)
	}

	var query = req.URL.Query()

	for _, data := range strings.Split(t.mapToGetParams(data, false), "&"){
		var parts = strings.Split(data, "=")
		var key string
		var value string
		if len(parts) == 2 {
			key = parts[0]
			value = parts[1]
		}else if len(parts) == 1 {
			key = parts[0]
		}
		query.Add(key, value)
	}

	return query.Encode()
}

func (t *transform)mapToGetParams(data interface{}, brackets bool)string{
	var result string
	if data != nil {
		var params = make([]string, 0)
		switch reflect.TypeOf(data).Kind() {
		case reflect.Float64:
			return fmt.Sprintf("=%v", data)
		case reflect.Int:
			return fmt.Sprintf("=%v", data)
		case reflect.Map:

			var str = make([]string, 0)
			for key, value := range data.(map[string]interface{}){
				for _, tmp := range strings.Split(t.mapToGetParams(value, true), "&"){
					if brackets {
						str = append(str, fmt.Sprintf("[%s]%s", key, tmp))
					}else{
						str = append(str, fmt.Sprintf("%s%s", key, tmp))
					}
				}
			}
			params = append(params, str...)
		case reflect.Slice:
			var str = make([]string, 0)
			var resultSlice = reflect.ValueOf(data)

			for i := 0; i < resultSlice.Len(); i++{
				for _, tmp := range strings.Split(t.mapToGetParams(resultSlice.Index(i).Interface(), true), "&"){
					str = append(str, fmt.Sprintf("[]%s", tmp))
				}
			}
			params = append(params, str...)
		case reflect.String:
			return fmt.Sprintf("=%s", data)
		case reflect.Struct:
			return fmt.Sprintf("=%v", data)
		default:
			return fmt.Sprintf("=%v", data)
		}
		result = strings.Join(params, "&")
	}

	return result
}

func (t *transform)ByteToBase64(data []byte)string{
	return base64.StdEncoding.EncodeToString(data)
}

func (t *transform)Base64ToByte(data string)[]byte{
	var result, err = base64.StdEncoding.DecodeString(data)
	if err != nil {
		panic(err)
	}
	return result
}