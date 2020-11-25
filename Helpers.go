package ulog

import (
	"bufio"
	"bytes"
)

// Logs the error to ERROR-level if it is not nil
func LogIfError(err error) {
	if err != nil {
		Error(err)
	}
}

// Logs the error to PANIC-level (and panicking after) if is not nil
func PanicIfError(err error) {
	if err != nil {
		Panic(err)
	}
}

// Logs if error received as second argument to ERROR-level is not nil (first argument is discarded)
func LogIfErrorSecondArg(_ interface{}, err error) {
	if err != nil {
		Error(err)
	}
}

// Logs if error received as second argument to PANIC-level (and panicking after) is not nil (first argument is discarded)
func PanicIfErrorSecondArg(_ interface{}, err error) {
	if err != nil {
		Panic(err)
	}
}

// Logs the error to INFO-level if it is not nil
func LogIfErrorToInfo(err error) {
	if err != nil {
		Info(err)
	}
}

// Logs if error received as second argument to INFO-level is not nil (first argument is discarded)
func LogIfErrorToInfoSecondArg(_ interface{}, err error) {
	if err != nil {
		Info(err)
	}
}

// Logs a byteArray line by line to Error
func LogByteArrayLineByLineToError(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, Errorf, prefix...)
}

// Logs a byteArray line by line to Warn
func LogByteArrayLineByLineToWarn(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, Warnf, prefix...)
}

// Logs a byteArray line by line to Info
func LogByteArrayLineByLineToInfo(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, Errorf, prefix...)
}

// Logs a byteArray line by line to Info
func LogByteArrayLineByLineToDebug(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, Debugf, prefix...)
}

// Logs a byteArray line by line to Trace
func LogByteArrayLineByLineToTrace(in []byte, prefix ...string) {
	LogByteArrayLineByLine(in, Tracef, prefix...)
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
