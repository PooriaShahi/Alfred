package helpers

import (
	"fmt"
	"reflect"
	"strings"
)

func Walk(m map[string]interface{}) {
	for k, v := range m {

		kind := reflect.ValueOf(v).Kind()

		if kind == reflect.Map {
			new := make(map[string]interface{})
			newValue := v.(map[string]interface{})
			for key, value := range newValue {
				new[k+"_"+key] = value
			}
			Walk(new)
		}

		if kind == reflect.Int {
			fmt.Printf("%v=%v\n", strings.ToUpper(k), v)
		}

		if kind == reflect.String {
			fmt.Printf("%v=%v\n", strings.ToUpper(k), v)
		}

		if kind == reflect.Bool {
			fmt.Printf("%v=%v\n", strings.ToUpper(k), v)
		}
	}
}
