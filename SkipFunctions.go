package ulog

import "sort"

var skipFunctions []string

func SkipFunctions() []string {
	return skipFunctions
}

func AddSkipFunctions(_skipFunctions ...string) {
	defaultSkipFunctions := []string{
		"github.com/dunv/ulog.LogIfError",
		"github.com/dunv/ulog.LogIfErrorSecondArg",
		"github.com/dunv/ulog.LogIfErrorToInfo",
		"github.com/dunv/ulog.LogIfErrorToInfoSecondArg",
	}
	desiredSkipFunctions := append(defaultSkipFunctions, _skipFunctions...)
	skipFunctionsToBeAdded := []string{}
	for index, desiredSkipFunction := range desiredSkipFunctions {
		i := sort.Search(len(skipFunctions), func(i int) bool { return skipFunctions[i] >= desiredSkipFunction })
		if i >= len(skipFunctions) {
			skipFunctionsToBeAdded = append(skipFunctionsToBeAdded, desiredSkipFunctions[index])
		}
	}
	sort.Strings(skipFunctionsToBeAdded)
}
