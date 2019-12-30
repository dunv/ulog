package ulog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

// Make sure the correct logging function is used. Formatting symbols should only be regarded when using a formatting function
func TestPrintOut(t *testing.T) {
	ResetFormat()
	buffer := setup(LEVEL_TRACE)
	msg := "halloWelt%s"
	Trace(msg)
	checkPrintOut(buffer, _LEVEL_TRACE, msg, t)

	buffer = setup(LEVEL_TRACE)
	Tracef("test %s", msg)
	checkPrintOut(buffer, _LEVEL_TRACE, fmt.Sprintf("test %s", msg), t)
}

func checkPrintOut(buffer *bytes.Buffer, level logLevelString, msg string, t *testing.T) {
	res, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Error(err)
	}

	var re = regexp.MustCompile(fmt.Sprintf(`(?m)^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d\.\d\d\d \| %s \| github.com\/dunv\/ulog printOut_test\.go:\d+ \(TestPrintOut\) \| %s$`, level, msg))
	if !re.Match(res) {
		t.Errorf(`did not log the correct output actual: "%s"`, string(res))
	}
}
