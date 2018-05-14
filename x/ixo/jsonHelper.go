package ixo

import (
	json "encoding/json"
	"fmt"
)

type JsonString struct {
	Value string
}

type JsonObject struct {
	Value map[string]interface{}
}

func (jo *JsonObject) String() string {
	output, err := json.MarshalIndent(jo.Value, "", "  ")
	if err != nil {
		panic(err)
	}
	return fmt.Sprintf("%v", string(output))
}
func (js *JsonString) ParseJSON() JsonObject {
	jsonBytes := []byte(js.Value)
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
