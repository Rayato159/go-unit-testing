package myTests

import "encoding/json"

func CompressToJSON(obj any) string {
	result, _ := json.Marshal(&obj)
	return string(result)
}
