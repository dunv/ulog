package ulog

import "strings"

func LogIfError(err error) {
	if err != nil {
		Error(err)
	}
}

func LogIfErrorSecondArg(_ interface{}, err error) {
	if err != nil {
		Error(err)
	}
}

func LogIfErrorToInfo(err error) {
	if err != nil {
		Info(err)
	}
}

func LogIfErrorToInfoSecondArg(_ interface{}, err error) {
	if err != nil {
		Info(err)
	}
}

func LogLevelFromString(levelRaw string) LogLevel {
	switch strings.ToUpper(levelRaw) {
	case "TRACE":
		return LEVEL_TRACE
	case "DEBUG":
		return LEVEL_DEBUG
	case "INFO":
		fallthrough
	case "INFORMATION":
		return LEVEL_INFO
	case "WARN":
		fallthrough
	case "WARNING":
		return LEVEL_WARNING
	case "ERROR":
		return LEVEL_ERROR
	case "FATAL":
		return LEVEL_FATAL
	default:
		Errorf("could not parse logLevel from string %s", levelRaw)
		return LEVEL_INFO
	}
}
