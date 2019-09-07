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
