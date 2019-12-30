package ulog

import (
	"errors"
	"testing"
)

func TestLogIfError(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	LogIfError(errors.New("testError"))
	checkDisregardPackage(buffer, _LEVEL_ERROR, "testError", t)
}

func TestPanicIfError(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {

			checkDisregardPackage(buffer, _LEVEL_ERROR, "testError", t)
		}
	}()

	PanicIfError(errors.New("testError"))
}

func TestLogIfErrorSecondArg(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	LogIfErrorSecondArg("disregard", errors.New("testError"))
	checkDisregardPackage(buffer, _LEVEL_ERROR, "testError", t)
}

func TestPanicIfErrorSecondArg(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("The code did not panic")
		} else {

			checkDisregardPackage(buffer, _LEVEL_ERROR, "testError", t)
		}
	}()

	PanicIfErrorSecondArg("disregard", errors.New("testError"))
}

func TestLogIfErrorToInfo(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	LogIfErrorToInfo(errors.New("testError"))
	checkDisregardPackage(buffer, _LEVEL_INFO, "testError", t)
}

func TestLogIfErrorToInfoSecondArg(t *testing.T) {
	buffer := setup(LEVEL_TRACE)
	LogIfErrorToInfoSecondArg("disregard", errors.New("testError"))
	checkDisregardPackage(buffer, _LEVEL_INFO, "testError", t)
}
