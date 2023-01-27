package ulog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestLoggingModifiedTemplate(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	SetFormatString("{{ .Time }} | {{ .Level }} | n/a | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n")
	msg := "halloWelt"
	Trace(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_TRACE, msg, t)

	buffer = setup(LEVEL_TRACE)
	SetFormatString("{{ .Time }} | {{ .Level }} | n/a | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n")
	Tracef("test %s", msg)
	checkWithModifiedTemplate(buffer, _LEVEL_TRACE, fmt.Sprintf("test %s", msg), t)

	buffer = setup(LEVEL_TRACE)
	SetFormatString("{{ .Time }} | {{ .Level }} | n/a | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n")
	Debug(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_DEBUG, msg, t)

	buffer = setup(LEVEL_TRACE)
	SetFormatString("{{ .Time }} | {{ .Level }} | n/a | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n")
	Info(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_INFO, msg, t)

	buffer = setup(LEVEL_TRACE)
	SetFormatString("{{ .Time }} | {{ .Level }} | n/a | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n")
	Warn(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_WARNING, msg, t)

	buffer = setup(LEVEL_TRACE)
	SetFormatString("{{ .Time }} | {{ .Level }} | n/a | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n")
	Error(msg)
	checkWithModifiedTemplate(buffer, _LEVEL_ERROR, msg, t)

}

func TestLoggingModifiedTime(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	msg := "halloWelt"
	SetTimestampFormat("2006-01-02 15:04:05")
	Trace(msg)
	checkWithModifiedTime(buffer, _LEVEL_TRACE, msg, t)
}

func checkWithModifiedTemplate(buffer *bytes.Buffer, level logLevelString, msg string, t *testing.T) {
	res, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Error(err)
	}

	regexString := fmt.Sprintf(`(?m)^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d\.\d\d\d \| %s \| n\/a \| github.com\/dunv\/ulog ulog\/Format_test\.go:\d+ \(TestLoggingModifiedTemplate\) \| %s$`, level, msg)
	re := regexp.MustCompile(regexString)

	if !re.Match(res) {
		t.Errorf(`did not log the correct output actual: "%s"`, string(res))
	}
}

func checkWithModifiedTime(buffer *bytes.Buffer, level logLevelString, msg string, t *testing.T) {
	res, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Error(err)
	}

	var re = regexp.MustCompile(fmt.Sprintf(`(?m)^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d \| %s \| github.com\/dunv\/ulog ulog\/Format_test\.go:\d+ \(TestLoggingModifiedTime\) \| %s$`, level, msg))
	if !re.Match(res) {
		t.Errorf(`did not log the correct output actual: "%s"`, string(res))
	}
}
