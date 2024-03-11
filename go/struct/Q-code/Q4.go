package qcode

import (
	"fmt"
	"reflect"
)

type Student struct {
	Name  string  `json:"name`
	Age   int     `json:"age`
	Score float64 `jason:score`
}

// func RefaectStructFunc() {
// 	reflect.TypeOf(Structest)
// }

func printStructInfo(s interface{}) {
	val := reflect.ValueOf(s)
	typ := val.Type()

	for i := 0; i < val.NumField(); i++ {
		field := typ.Field(i)
		value := val.Field(i).Interface()
		tag := field.Tag.Get("json")
		fmt.Printf("%s (%s): %v\n", field.Name, tag, value)
	}
}

func Q4() {
}
