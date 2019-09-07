package ulog

import (
	"sort"
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
}

func SkipFunctions() []string {
	return skipFunctions
}

func AddSkipFunctions(_skipFunctions ...string) {
	sort.Strings(skipFunctions)
	skipFunctionsToBeAdded := []string{}
	for index, desiredSkipFunction := range _skipFunctions {
		i := sort.SearchStrings(skipFunctions, desiredSkipFunction)
		if !(i == 0 && skipFunctions[i] == desiredSkipFunction) && !(i > 0 && i < len(skipFunctions)) {
			skipFunctionsToBeAdded = append(skipFunctionsToBeAdded, _skipFunctions[index])
		}
	}
	skipFunctions = append(skipFunctions, skipFunctionsToBeAdded...)
}
