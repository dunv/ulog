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

func TestGetReplaceFunctions(t *testing.T) {
	AddReplaceFunctions(map[string]string{
		"github.com/dunv/ulog.TestReplaceFunctions": "myReplacement",
	})

	fns := ReplaceFunctions()
	if fn, ok := fns["github.com/dunv/ulog.TestReplaceFunctions"]; !ok {
		t.Errorf("did not store function correctly")
	} else {
		if fn != "myReplacement" {
			t.Errorf("did not store function correctly")
		}
	}
}
