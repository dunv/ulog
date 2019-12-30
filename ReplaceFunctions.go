package ulog

import (
	"github.com/dunv/uhelpers"
)

var replaceFunctions map[string]string = map[string]string{}

// Get replace-functions
func ReplaceFunctions() map[string]string {
	return replaceFunctions
}

// Add a single replace function. This is used to make log-output origin easier to identify.
// Instead of getting a filename, package and function, a desired identifier can be chosen.
// for example
//  ulog.AddReplaceFunction("github.com/dunv/uhttp/middlewares.AddLogging.func1", "uhttp.Logging")
//  ulog.AddReplaceFunction("github.com/dunv/uhttp.Handle", "uhttp.Handle")
func AddReplaceFunction(fnToBeReplaced string, replaceWithFn string) {
	keys := uhelpers.StringKeysFromMap(replaceFunctions)
	if !uhelpers.SliceContainsItem(keys, fnToBeReplaced) {
		replaceFunctions[fnToBeReplaced] = replaceWithFn
	}
}

// Add a multiple replace functions. See AddReplaceFunction for more detail
func AddReplaceFunctions(m map[string]string) {
	for key := range m {
		AddReplaceFunction(key, m[key])
	}
}
