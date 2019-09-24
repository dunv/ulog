package ulog

import (
	"fmt"
	"reflect"
	"strings"
)

func LogEnvStruct(envStruct interface{}, prefix string) {
	LogTaggedStructWithMaskingAndWarning(envStruct, "env", "mask", "warnIf", prefix)
}

func LogTaggedStructWithMaskingAndWarning(taggedStruct interface{}, tag string, maskTag string, warnTag string, prefix string) {
	typeOf := reflect.TypeOf(taggedStruct)
	valueOf := reflect.ValueOf(taggedStruct)

	allWarnings := []string{}
	allOutput := []string{}

	for i := 0; i < typeOf.NumField(); i++ {
		tag := typeOf.Field(i).Tag.Get(tag)
		maskTag := typeOf.Field(i).Tag.Get(maskTag)
		warnTag := typeOf.Field(i).Tag.Get(warnTag)
		if tag != "" {
			if maskTag != "" {
				allOutput = append(allOutput, fmt.Sprintf("%s%s = %v", prefix, tag, "***"))
			} else {
				allOutput = append(allOutput, fmt.Sprintf("%s%s = %v", prefix, tag, getStringFromReflect(valueOf.Field(i))))
			}

			if warnTag == getStringFromReflect(valueOf.Field(i)) {
				allWarnings = append(allWarnings, fmt.Sprintf("%s%s has warning value", prefix, tag))
			}
		}
	}

	if len(allWarnings) > 0 {
		Warn("WARNINGS ..............................................................")
	}
	for _, warning := range allWarnings {
		Warn(warning)
	}
	if len(allWarnings) > 0 {
		Warn("Config ................................................................")
	}
	for _, output := range allOutput {
		Info(output)
	}
}

func getStringFromReflect(val reflect.Value) string {
	switch val.Interface().(type) {
	case string:
		return strings.ReplaceAll(val.Interface().(string), "%", "%%")
	default:
		return fmt.Sprintf("%v", val)
	}
}
