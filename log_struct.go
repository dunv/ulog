package ulog

import (
	"reflect"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// Log out a struct line by line to Debug
func DebugStruct(taggedStruct any, prefix string) {
	logStruct(zapcore.DebugLevel, taggedStruct, prefix)
}

// Log out a struct line by line to Info
func InfoStruct(taggedStruct any, prefix string) {
	logStruct(zapcore.InfoLevel, taggedStruct, prefix)
}

// Log out a struct line by line to Warn
func WarnStruct(taggedStruct any, prefix string) {
	logStruct(zapcore.WarnLevel, taggedStruct, prefix)
}

// Log out a struct line by line to Error
func ErrorStruct(taggedStruct any, prefix string) {
	logStruct(zapcore.ErrorLevel, taggedStruct, prefix)
}

// Log out a struct line by line
func logStruct(lvl zapcore.Level, taggedStruct any, prefix string) {
	usedTaggedStruct := taggedStruct
	if reflect.ValueOf(taggedStruct).Kind() == reflect.Ptr {
		usedTaggedStruct = reflect.ValueOf(taggedStruct).Elem().Interface()
	}

	typeOf := reflect.TypeOf(usedTaggedStruct)
	valueOf := reflect.ValueOf(usedTaggedStruct)
	logger := skipOneSugaredLogger.WithOptions(zap.AddCallerSkip(1))

	for i := 0; i < valueOf.NumField(); i++ {
		derefFieldVal := valueOf.Field(i)
		derefFieldName := typeOf.Field(i).Name
		logger.Logf(lvl, "%s%s = %v", prefix, derefFieldName, getStringFromReflect(derefFieldVal))
	}
}
