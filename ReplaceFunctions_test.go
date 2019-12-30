package ulog

import (
	"testing"
)

func TestReplaceFunctions(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	AddReplaceFunction("github.com/dunv/ulog.TestReplaceFunctions", "myReplacement")
	Trace("halloWelt")
	checkCustomFunction(buffer, _LEVEL_TRACE, "halloWelt", "myReplacement", t)
}
