package utils

import "github.com/tidwall/gjson"
import "fmt"

const json = `{"name":{"first":"Janet","last":"Prichard"},"age":47}`

func ToJSONString() {
	value := gjson.Get(json, "name.last")
	fmt.Println(value.String())
}
