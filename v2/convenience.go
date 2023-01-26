package v2

import (
	"bufio"
	"bytes"
)

// Logs the error to ERROR-level if it is not nil
func LogIfError(err error) {
	if err != nil {
		skipOneLogger.Error(err)
	}
}

// Logs the error to TRACE-level if is not nil
func TraceIfError(err error) {
	if err != nil {
		skipOneLogger.Trace(err)
	}
}

// Logs the error to DEBUG-level if is not nil
func DebugIfError(err error) {
	if err != nil {
		skipOneLogger.Debug(err)
	}
}

// Logs the error to INFO-level if is not nil
func InfoIfError(err error) {
	if err != nil {
		skipOneLogger.Info(err)
	}
}

// Logs the error to WARN-level if is not nil
func WarnIfError(err error) {
	if err != nil {
		skipOneLogger.Warn(err)
	}
}

// Logs the error to FATAL-level if is not nil
func FatalIfError(err error) {
	if err != nil {
		skipOneLogger.Fatal(err)
	}
}

// Logs the error to PANIC-level (and panicking after) if is not nil
func PanicIfError(err error) {
	if err != nil {
		skipOneLogger.Panic(err)
	}
}

// Logs if error received as second argument to ERROR-level is not nil (first argument is discarded)
func LogIfErrorSecondArg(input interface{}, err error) interface{} {
	if err != nil {
		skipOneLogger.Error(err)
		return input
	}
	return input
}

// Logs if error received as second argument to FATAL-level (and panicking after) is not nil (first argument is discarded)
func FatalIfErrorSecondArg(input interface{}, err error) interface{} {
	if err != nil {
		skipOneLogger.Fatal(err)
		return input
	}
	return input
}

// Logs if error received as second argument to PANIC-level (and panicking after) is not nil (first argument is discarded)
func PanicIfErrorSecondArg(input interface{}, err error) interface{} {
	if err != nil {
		skipOneLogger.Panic(err)
		return input
	}
	return input
}

// Logs the error to INFO-level if it is not nil
func LogIfErrorToInfo(err error) {
	if err != nil {
		skipOneLogger.Info(err)
	}
}

// Logs if error received as second argument to INFO-level is not nil (first argument is discarded)
func LogIfErrorToInfoSecondArg(_ interface{}, err error) {
	if err != nil {
		skipOneLogger.Info(err)
	}
}

// Logs a byteArray line by line to Error
func LogByteArrayLineByLineToError(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, skipOneLogger.Errorf, prefix...)
}

// Logs a byteArray line by line to Warn
func LogByteArrayLineByLineToWarn(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, skipOneLogger.Warnf, prefix...)
}

// Logs a byteArray line by line to Info
func LogByteArrayLineByLineToInfo(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, skipOneLogger.Infof, prefix...)
}

// Logs a byteArray line by line to Debug
func LogByteArrayLineByLineToDebug(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, skipOneLogger.Debugf, prefix...)
}

// Logs a byteArray line by line
func LogByteArrayLineByLine(in []byte, f func(string, ...interface{}), prefix ...string) {
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
