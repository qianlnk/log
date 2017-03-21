package log

import (
	"reflect"
)

func Map(val map[string]string) Fields {
	tmp := make(map[string]interface{})
	for k, v := range val {
		tmp[k] = v
	}
	return Fields(tmp)
}

func Struct(st interface{}) Fields {
	val := reflect.ValueOf(st)
	typ := val.Type()

	fields := make(map[string]interface{})
	for i := 0; i < val.NumField(); i++ {
		fields[typ.Field(i).Name] = val.Field(i).Interface()
	}

	return fields
}
