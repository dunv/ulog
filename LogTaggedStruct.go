package ulog

import (
	"reflect"
)

func LogEnvStruct(envStruct interface{}) {
	LogTaggedStruct(envStruct, "env")
}

func LogTaggedStruct(taggedStruct interface{}, tag string) {
	typeOf := reflect.TypeOf(taggedStruct)
	valueOf := reflect.ValueOf(taggedStruct)
	for i := 0; i < typeOf.NumField(); i++ {
		tag := typeOf.Field(i).Tag.Get(tag)
		if tag != "" {
			Infof("%s = %v", tag, valueOf.Field(i))
		}
	}
}
