package debug

import (
	"encoding/json"
	"fmt"
)

func PrintStruct(data interface{}) {
	dataJson, _ := json.MarshalIndent(data, "", "\t")
	fmt.Println(string(dataJson))
}
