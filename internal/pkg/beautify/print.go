package beautify

import (
	"encoding/json"
	"fmt"
	"log"
)

func JSONPrint(v interface{}) {
	fmt.Println(JSONString(v))
}

func JSONString(v interface{}) string {
	b, err := json.MarshalIndent(v, "", "  ")
	if err != nil {
		log.Panicf("JSONString: %s", err.Error())
	}
	return string(b)
}
