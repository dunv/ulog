package ulog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestLogging(t *testing.T) {
	ResetFormat()

	buffer := setup(LEVEL_TRACE)
	msg := "halloWelt"
	Trace(msg)
	check(buffer, _LEVEL_TRACE, msg, t)

	buffer = setup(LEVEL_TRACE)
	Tracef("test %s", msg)
	check(buffer, _LEVEL_TRACE, fmt.Sprintf("test %s", msg), t)

	buffer = setup(LEVEL_TRACE)
	Debug(msg)
	check(buffer, _LEVEL_DEBUG, msg, t)

	buffer = setup(LEVEL_TRACE)
	Info(msg)
	check(buffer, _LEVEL_INFO, msg, t)

	buffer = setup(LEVEL_TRACE)
	Warn(msg)
	check(buffer, _LEVEL_WARNING, msg, t)

	buffer = setup(LEVEL_TRACE)
	Error(msg)
	check(buffer, _LEVEL_ERROR, msg, t)

}

func setup(logLevel LogLevel) *bytes.Buffer {
	tmp := []byte{}
	buffer := bytes.NewBuffer(tmp)
	SetLogLevel(logLevel)
	SetWriter(buffer, nil)
	return buffer
}

func check(buffer *bytes.Buffer, level logLevelString, msg string, t *testing.T) {
	res, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Error(err)
	}

	var re = regexp.MustCompile(fmt.Sprintf(`(?m)^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d\.\d\d\d \| %s \| github.com\/dunv\/ulog ULogger_test\.go:\d+ \(TestLogging\) \| %s$`, level, msg))
	if !re.Match(res) {
		t.Errorf(`did not log the correct output actual: "%s"`, string(res))
	}
}
