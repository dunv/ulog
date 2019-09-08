package ulog

import (
	"github.com/dunv/uhelpers"
)

var replaceFunctions map[string]string = map[string]string{}

func ReplaceFunctions() map[string]string {
	return replaceFunctions
}

func AddReplaceFunction(fnToBeReplaced string, replaceWithFn string) {
	keys := uhelpers.StringKeysFromMap(replaceFunctions)
	if !uhelpers.SliceContainsItem(keys, fnToBeReplaced) {
		replaceFunctions[fnToBeReplaced] = replaceWithFn
	}
}

func AddReplaceFunctions(m map[string]string) {
	for key := range m {
		AddReplaceFunction(key, m[key])
	}
}
