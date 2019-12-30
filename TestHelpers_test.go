package ulog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func setup(logLevel LogLevel) *bytes.Buffer {
	ResetFormat()
	ResetReplaceFunctions()
	UnsetDebug()
	tmp := []byte{}
	buffer := bytes.NewBuffer(tmp)
	SetLogLevel(logLevel)
	SetWriter(buffer, nil)
	return buffer
}

func checkDisregardPackage(buffer *bytes.Buffer, level logLevelString, msg string, t *testing.T) {
	res, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Error(err)
	}

	var re = regexp.MustCompile(fmt.Sprintf(`(?m)^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d\.\d\d\d \| %s \| github.com\/dunv\/ulog[\.a-zA-Z]* [a-zA-Z_]+\.go:\d+ \([a-zA-Z_\./]+\) \| %s$`, level, msg))
	if !re.Match(res) {
		t.Errorf(`did not log the correct output actual: "%s"`, string(res))
	}
}

func checkCustomFunction(buffer *bytes.Buffer, level logLevelString, msg string, function string, t *testing.T) {
	res, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Error(err)
	}

	var re = regexp.MustCompile(fmt.Sprintf(`(?m)^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d\.\d\d\d \| %s \| %s \| %s$`, level, function, msg))
	if !re.Match(res) {
		t.Errorf(`did not log the correct output actual: "%s"`, string(res))
	}
}
