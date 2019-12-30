package ulog

import (
	"sort"

	"github.com/dunv/uhelpers"
)

var skipFunctions []string = []string{
	"github.com/dunv/ulog.LogIfError",
	"github.com/dunv/ulog.LogIfErrorSecondArg",
	"github.com/dunv/ulog.PanicIfError",
	"github.com/dunv/ulog.PanicIfErrorSecondArg",
	"github.com/dunv/ulog.LogIfErrorToInfo",
	"github.com/dunv/ulog.LogIfErrorToInfoSecondArg",
	"github.com/dunv/ulog.ULog.Trace",
	"github.com/dunv/ulog.ULog.Tracef",
	"github.com/dunv/ulog.ULog.Debug",
	"github.com/dunv/ulog.ULog.Debugf",
	"github.com/dunv/ulog.ULog.Info",
	"github.com/dunv/ulog.ULog.Infof",
	"github.com/dunv/ulog.ULog.Warn",
	"github.com/dunv/ulog.ULog.Warnf",
	"github.com/dunv/ulog.ULog.Error",
	"github.com/dunv/ulog.ULog.Errorf",
	"github.com/dunv/ulog.ULog.Panic",
	"github.com/dunv/ulog.ULog.Panicf",
	"github.com/dunv/ulog.ULog.Fatal",
	"github.com/dunv/ulog.ULog.Fatalf",
	"github.com/dunv/ulog.LogTaggedStruct",
	"github.com/dunv/ulog.LogTaggedStructWithMaskingAndWarning",
	"github.com/dunv/ulog.LogEnvStruct",
}

// Get currently configured skipFunctions
func SkipFunctions() []string {
	return skipFunctions
}

// // Make expected output (which is only for info, not for debugging) more readable
// Add skipFunctions. These functions will be skipped, when trying to determine the origin of a log-line.
// This way helper functions which are in use for multiple different origins can be omitted.
// For example
//  ulog.AddSkipFunctions(
//  	"github.com/dunv/uhttp.RenderError",
//  	"github.com/dunv/uhttp/helpers.RenderError",
//  	"github.com/dunv/uhttp.RenderErrorWithStatusCode",
//  	"github.com/dunv/uhttp/helpers.RenderErrorWithStatusCode",
//  	"github.com/dunv/uhttp.RenderMessage",
//  	"github.com/dunv/uhttp/helpers.RenderMessage",
//  	"github.com/dunv/uhttp.RenderMessageWithStatusCode",
//  	"github.com/dunv/uhttp/helpers.RenderMessageWithStatusCode",
//  	"github.com/dunv/uhttp.renderMessageWithStatusCode",
//  	"github.com/dunv/uhttp/helpers.renderMessageWithStatusCode",
//  	"github.com/dunv/uhttp/helpers.renderErrorWithStatusCode",
//  	"github.com/dunv/uhttp.renderErrorWithStatusCode",
//  )
func AddSkipFunctions(_skipFunctions ...string) {
	sort.Strings(skipFunctions)
	skipFunctionsToBeAdded := []string{}

	for index, desiredSkipFunction := range _skipFunctions {
		if desiredSkipFunction != "main.main" {
			if !uhelpers.SliceContainsItem(skipFunctions, desiredSkipFunction) {
				skipFunctionsToBeAdded = append(skipFunctionsToBeAdded, _skipFunctions[index])
			}
		} else {
			Error("cannot add main.main as skip function")
		}
	}

	skipFunctions = append(skipFunctions, skipFunctionsToBeAdded...)
}
