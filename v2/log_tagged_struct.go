package v2

import (
	"fmt"
	"reflect"
	"strings"

	"go.uber.org/zap"
)

// Helper function to log out a struct annotated with "env", "mask" and "warnIf" annotations
// Good to use with the "github.com/codingconcepts/env" package
func LogEnvStruct(envStruct interface{}, prefix string) {
	LogTaggedStructWithMaskingAndWarning(envStruct, "env", "mask", "warnIf", prefix, true, skipOneSugaredLogger.WithOptions(zap.AddCallerSkip(1)))
}

// Helper function to log out a struct annotated with "json", "mask" and "warnIf" annotations
// Good to use with the "encoding/json" package
func LogJSONStruct(envStruct interface{}, prefix string) {
	LogTaggedStructWithMaskingAndWarning(envStruct, "json", "mask", "warnIf", prefix, false, skipOneSugaredLogger.WithOptions(zap.AddCallerSkip(1)))
}

// Helper function to log out a struct
// should only be used if LogEnvStruct() or LogJSONStruct() do not meet your requirements
func LogTaggedStructWithMaskingAndWarning(taggedStruct interface{}, tag string, maskTag string, warnTag string, prefix string, printHeaders bool, logger *zap.SugaredLogger) {
	// Dereference if needed
	usedTaggedStruct := taggedStruct
	if reflect.ValueOf(taggedStruct).Kind() == reflect.Ptr {
		usedTaggedStruct = reflect.ValueOf(taggedStruct).Elem().Interface()
	}

	typeOf := reflect.TypeOf(usedTaggedStruct)
	valueOf := reflect.ValueOf(usedTaggedStruct)

	allWarnings := []string{}
	allOutput := []string{}

	for i := 0; i < valueOf.NumField(); i++ {
		derefFieldVal := valueOf.Field(i)
		derefFieldType := typeOf.Field(i)

		tag := derefFieldType.Tag.Get(tag)
		maskTag := derefFieldType.Tag.Get(maskTag)
		warnTag := derefFieldType.Tag.Get(warnTag)
		if tag != "" {
			if maskTag != "" {
				allOutput = append(allOutput, fmt.Sprintf("%s%s = %v", prefix, tag, "***"))
			} else {
				allOutput = append(allOutput, fmt.Sprintf("%s%s = %v", prefix, tag, getStringFromReflect(derefFieldVal)))
			}

			if warnTag == getStringFromReflect(valueOf.Field(i)) {
				allWarnings = append(allWarnings, fmt.Sprintf("%s%s has warning value", prefix, tag))
			}
		}
	}

	if printHeaders && len(allWarnings) > 0 {
		logger.Warn("WARNINGS ..............................................................")
	}
	for _, warning := range allWarnings {
		logger.Warn(warning)
	}
	if printHeaders && len(allOutput) > 0 {
		logger.Info("Config ................................................................")
	}
	for _, output := range allOutput {
		logger.Info(output)
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