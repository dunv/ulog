package ulog

import (
	"sort"

	"github.com/dunv/uhelpers"
)

var skipFunctions []string = []string{
	"github.com/dunv/ulog.LogIfError",
	"github.com/dunv/ulog.LogIfErrorSecondArg",
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
	"github.com/dunv/ulog.LogEnvStruct",
}

func SkipFunctions() []string {
	return skipFunctions
}

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
