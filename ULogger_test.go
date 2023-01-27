package ulog

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"regexp"
	"testing"
)

func TestLoggingHelpers(t *testing.T) {

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
	Debugf(msg)
	check(buffer, _LEVEL_DEBUG, msg, t)

	buffer = setup(LEVEL_TRACE)
	Info(msg)
	check(buffer, _LEVEL_INFO, msg, t)

	buffer = setup(LEVEL_TRACE)
	Infof(msg)
	check(buffer, _LEVEL_INFO, msg, t)

	buffer = setup(LEVEL_TRACE)
	Warn(msg)
	check(buffer, _LEVEL_WARNING, msg, t)

	buffer = setup(LEVEL_TRACE)
	Warnf(msg)
	check(buffer, _LEVEL_WARNING, msg, t)

	buffer = setup(LEVEL_TRACE)
	Error(msg)
	check(buffer, _LEVEL_ERROR, msg, t)

	buffer = setup(LEVEL_TRACE)
	Errorf(msg)
	check(buffer, _LEVEL_ERROR, msg, t)

	buffer = setup(LEVEL_TRACE)
	LogWithLevel(LEVEL_TRACE, msg)
	check(buffer, _LEVEL_TRACE, msg, t)

	buffer = setup(LEVEL_TRACE)
	LogWithLevel(LEVEL_DEBUG, msg)
	check(buffer, _LEVEL_DEBUG, msg, t)

	buffer = setup(LEVEL_TRACE)
	LogWithLevel(LEVEL_INFO, msg)
	check(buffer, _LEVEL_INFO, msg, t)

	buffer = setup(LEVEL_TRACE)
	LogWithLevel(LEVEL_WARNING, msg)
	check(buffer, _LEVEL_WARNING, msg, t)

	buffer = setup(LEVEL_TRACE)
	LogWithLevel(LEVEL_ERROR, msg)
	check(buffer, _LEVEL_ERROR, msg, t)
}

func TestLoggingPanic(t *testing.T) {
	buffer := setup(LEVEL_TRACE)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {

			checkDisregardPackage(buffer, _LEVEL_ERROR, "testError", t)
		}
	}()

	Panic("testError")
}

func TestLoggingPanicf(t *testing.T) {
	buffer := setup(LEVEL_TRACE)

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {
			checkDisregardPackage(buffer, _LEVEL_ERROR, "testError test1 test2", t)
		}
	}()

	Panicf("testError %s %s", "test1", "test2")
}

func TestLoggingFatal(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	SetDebug() // required, otherwise we will exit(1)
	Fatal("testError")
	checkDisregardPackage(buffer, _LEVEL_FATAL, "testError", t)
}

func TestLoggingFatalf(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	SetDebug() // required, otherwise we will exit(1)
	Fatalf("testError %s", "test")
	checkDisregardPackage(buffer, _LEVEL_FATAL, "testError test", t)
}

func TestLoggingWithCustomLogger(t *testing.T) {
	tmpLogger := NewUlog()

	buffer := setup(LEVEL_TRACE)
	tmpLogger.Trace("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_TRACE, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Tracef("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_TRACE, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Debug("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_DEBUG, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Debugf("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_DEBUG, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Info("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_INFO, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Infof("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_INFO, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Warn("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_WARNING, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Warnf("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_WARNING, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Error("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_ERROR, "halloWelt", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger.Error("halloWelt")
	checkDisregardPackage(buffer, _LEVEL_ERROR, "halloWelt", t)

}

func TestLoggingPanicCustom(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	tmpLogger := NewUlog()

	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {

			checkDisregardPackage(buffer, _LEVEL_ERROR, "testError", t)
		}
	}()

	tmpLogger.Panic("testError")
}

func TestLoggingFatalCustom(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	tmpLogger := NewUlog()
	SetDebug() // otherwise the program would exit
	tmpLogger.Fatal("testError")
	checkDisregardPackage(buffer, _LEVEL_FATAL, "testError", t)
}

func TestLoggingFatalfCustom(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	tmpLogger := NewUlog()
	SetDebug() // otherwise the program would exit
	tmpLogger.Fatalf("testError %s", "test")
	checkDisregardPackage(buffer, _LEVEL_FATAL, "testError test", t)
}

// We had a variable-shadowing problem here before -> test struct-logger explicitly
func TestLoggingWithLevel(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	tmpLogger := NewUlog()
	tmpLogger.LogWithLevelf(LEVEL_TRACE, "testTrace")
	checkDisregardPackage(buffer, _LEVEL_TRACE, "testTrace", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger = NewUlog()
	tmpLogger.LogWithLevelf(LEVEL_DEBUG, "testDebug")
	checkDisregardPackage(buffer, _LEVEL_DEBUG, "testDebug", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger = NewUlog()
	tmpLogger.LogWithLevelf(LEVEL_INFO, "testInfo")
	checkDisregardPackage(buffer, _LEVEL_INFO, "testInfo", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger = NewUlog()
	tmpLogger.LogWithLevelf(LEVEL_WARNING, "testWarn")
	checkDisregardPackage(buffer, _LEVEL_WARNING, "testWarn", t)

	buffer = setup(LEVEL_TRACE)
	tmpLogger = NewUlog()
	tmpLogger.LogWithLevelf(LEVEL_ERROR, "testError")
	checkDisregardPackage(buffer, _LEVEL_ERROR, "testError", t)
}

func TestIfInterfaceIsMatchedByULog(t *testing.T) {
	var customLogger ULogger = NewUlog()
	_ = customLogger // circumvent compiler problem
}

func check(buffer *bytes.Buffer, level logLevelString, msg string, t *testing.T) {
	res, err := ioutil.ReadAll(buffer)
	if err != nil {
		t.Error(err)
	}

	var re = regexp.MustCompile(fmt.Sprintf(`(?m)^\d\d\d\d-\d\d-\d\d \d\d:\d\d:\d\d\.\d\d\d \| %s \| github.com\/dunv\/ulog ulog\/ULogger_test\.go:\d+ \(TestLoggingHelpers\) \| %s$`, level, msg))
	if !re.Match(res) {
		t.Errorf(`did not log the correct output actual: "%s"`, string(res))
	}
}
