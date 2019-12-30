package ulog

import (
	"fmt"
	"testing"
)

func TestAddSkipFunctions(t *testing.T) {

	before := len(SkipFunctions())
	AddSkipFunctions("a", "b", "c")
	after := len(SkipFunctions())

	if after-before != 3 {
		t.Error(fmt.Errorf("did not add a,b,c correctly"))
	}

	before = len(SkipFunctions())
	AddSkipFunctions("a")
	after = len(SkipFunctions())
	if after-before != 0 {
		t.Error(fmt.Errorf("did not add a correctly"))
	}

	before = len(SkipFunctions())
	AddSkipFunctions("b")
	after = len(SkipFunctions())
	if after-before != 0 {
		t.Error(fmt.Errorf("did not add b correctly"))
	}

	before = len(SkipFunctions())
	AddSkipFunctions("c")
	after = len(SkipFunctions())
	if after-before != 0 {
		t.Error(fmt.Errorf("did not add c correctly"))
	}

	before = len(SkipFunctions())
	AddSkipFunctions("a", "b", "c")
	after = len(SkipFunctions())
	if after-before != 0 {
		t.Error(fmt.Errorf("did not add c correctly"))
	}

}

func TestSkipFunctions(t *testing.T) {
	buffer := setup(LEVEL_TRACE)

	// Skip this function -> log should tell us testing.tRunner is the logging entity
	AddSkipFunctions("github.com/dunv/ulog.TestSkipFunctions")

	// Replacing is tested separately and is usable here (this way we do not need to pay attention if the line of code or the file of the testing std-lib changes)
	AddReplaceFunction("testing.tRunner", "replacedRunner")

	Trace("halloWelt")
	checkCustomFunction(buffer, _LEVEL_TRACE, "halloWelt", "replacedRunner", t)
}

func TestGetSkipFunctions(t *testing.T) {
	fns := SkipFunctions()
	if len(fns) != len(skipFunctions) {
		t.Errorf("could not retrieve skipFunctions")
	}
}
