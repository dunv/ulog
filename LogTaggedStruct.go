package ulog

import (
	"fmt"
	"reflect"
)

func LogEnvStruct(envStruct interface{}, prefix string) {
	LogTaggedStruct(envStruct, "env", prefix)
}

func LogTaggedStruct(taggedStruct interface{}, tag string, prefix string) {
	typeOf := reflect.TypeOf(taggedStruct)
	valueOf := reflect.ValueOf(taggedStruct)
	for i := 0; i < typeOf.NumField(); i++ {
		tag := typeOf.Field(i).Tag.Get(tag)
		if tag != "" {
			Infof(fmt.Sprintf("%s", prefix)+"%s = %v", tag, valueOf.Field(i))
		}
	}
}
