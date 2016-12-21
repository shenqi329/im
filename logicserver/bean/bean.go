package bean

import (
	"encoding/json"
)

func StructToJsonString(v interface{}) string {
	b, err := json.Marshal(v)

	if err != nil {
		return ""
	}
	return string(b)
}
