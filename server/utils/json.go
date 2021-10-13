package utils

import "encoding/json"

func ToObj(val string, obj interface{}) {
	err := json.Unmarshal([]byte(val), obj)
	if err != nil {
		panic(err)
	}
}

func ToStr(obj interface{}) string {
	bytes, err := json.Marshal(obj)
	if err != nil {
		panic(err)
	}
	return string(bytes)
}
