package ulog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestLoggingModifiedTemplate(t *testing.T) {
	SetFormatString("{{ .Time }} | {{ .Level }} | n/a | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n")

	buffer := setup(LEVEL_TRACE)
	msg := "halloWelt"
	Trace(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_TRACE, msg, t)

	buffer = setup(LEVEL_TRACE)
	Tracef("test %s", msg)
	checkWithModifiedTemplate(buffer, _LEVEL_TRACE, fmt.Sprintf("test %s", msg), t)

	buffer = setup(LEVEL_TRACE)
	Debug(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_DEBUG, msg, t)

	buffer = setup(LEVEL_TRACE)
	Info(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_INFO, msg, t)

	buffer = setup(LEVEL_TRACE)
	Warn(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_WARNING, msg, t)

	buffer = setup(LEVEL_TRACE)
	Error(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_ERROR, msg, t)

}

func checkWithModifiedTemplate(buffer *bytes.Buffer, level logLevelString, msg string, t *testing.T) {
	res, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Error(err)
	}

	var re = regexp.MustCompile(fmt.Sprintf(`(?m)^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d\.\d\d\d \| %s \| n\/a \| github.com\/dunv\/ulog Format_test\.go:\d+ \(TestLoggingModifiedTemplate\) \| %s$`, level, msg))
	if !re.Match(res) {
		t.Errorf(`did not log the correct output actual: "%s"`, string(res))
	}
}
