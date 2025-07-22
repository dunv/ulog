package ulog

import (
	"bufio"
	"bytes"
)

// Logs the error to ERROR-level if it is not nil
func LogIfError(err error) {
	if err != nil {
		SkipOneS().Error(err)
	}
}

// Logs the error to DEBUG-level if is not nil
func DebugIfError(err error) {
	if err != nil {
		SkipOneS().Debug(err)
	}
}

// Logs the error to INFO-level if is not nil
func InfoIfError(err error) {
	if err != nil {
		SkipOneS().Info(err)
	}
}

// Logs the error to WARN-level if is not nil
func WarnIfError(err error) {
	if err != nil {
		SkipOneS().Warn(err)
	}
}

// Logs the error to FATAL-level if is not nil
func FatalIfError(err error) {
	if err != nil {
		SkipOneS().Fatal(err)
	}
}

// Logs the error to PANIC-level (and panicking after) if is not nil
func PanicIfError(err error) {
	if err != nil {
		SkipOneS().Panic(err)
	}
}

// Logs if error received as second argument to ERROR-level is not nil (first argument is discarded)
func LogIfErrorSecondArg(input any, err error) any {
	if err != nil {
		SkipOneS().Error(err)
		return input
	}
	return input
}

// Logs if error received as second argument to FATAL-level (and panicking after) is not nil (first argument is discarded)
func FatalIfErrorSecondArg(input any, err error) any {
	if err != nil {
		SkipOneS().Fatal(err)
		return input
	}
	return input
}

// Logs if error received as second argument to PANIC-level (and panicking after) is not nil (first argument is discarded)
func PanicIfErrorSecondArg(input any, err error) any {
	if err != nil {
		SkipOneS().Panic(err)
		return input
	}
	return input
}

// Logs the error to INFO-level if it is not nil
func LogIfErrorToInfo(err error) {
	if err != nil {
		SkipOneS().Info(err)
	}
}

// Logs if error received as second argument to INFO-level is not nil (first argument is discarded)
func LogIfErrorToInfoSecondArg(_ any, err error) {
	if err != nil {
		SkipOneS().Info(err)
	}
}

// Logs a byteArray line by line to Error
func LogByteArrayLineByLineToError(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, SkipOneS().Errorf, prefix...)
}

// Logs a byteArray line by line to Warn
func LogByteArrayLineByLineToWarn(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, SkipOneS().Warnf, prefix...)
}

// Logs a byteArray line by line to Info
func LogByteArrayLineByLineToInfo(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, SkipOneS().Infof, prefix...)
}

// Logs a byteArray line by line to Debug
func LogByteArrayLineByLineToDebug(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, SkipOneS().Debugf, prefix...)
}

// Logs a byteArray line by line
func LogByteArrayLineByLine(in []byte, f func(string, ...any), prefix ...string) {
	bufReader := bufio.NewReader(bytes.NewReader(in))
	for {
		line, _, err := bufReader.ReadLine()
		if err != nil {
			break
		}

		if len(prefix) == 1 {
			f("%s%s", prefix[0], line)
		} else {
			f("%s", line)
		}
	}
}
