package ulog

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
