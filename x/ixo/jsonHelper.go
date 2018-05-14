package ixo

import (
	json "encoding/json"
	"fmt"
)

type JsonString struct {
	value string
}

type JsonObject struct {
	value map[string]interface{}
}

func (jo *JsonObject) String() string {
	output, err := json.MarshalIndent(jo.value, "", "  ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", string(output))
}
func (js *JsonString) ParseJSON() JsonObject {
	jsonBytes := []byte(js.value)
	var f interface{}
	err := json.Unmarshal(jsonBytes, &f)
	if err != nil {
		panic(err)
	}
	m := f.(map[string]interface{})
	return JsonObject{m}

}

/*
EXAMPLE:

func main() {

	b := JsonString{`{"Name":"Wednesday","Age":6,"Parents":[{"Name":"Gomez"},{"Name":"Morticia"}]}`}
	jo := b.ParseJSON()
	fmt.Printf(jo.String())

}
*/
