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
	case string(_LEVEL_TRACE):
		return LEVEL_TRACE
	case string(_LEVEL_DEBUG):
		return LEVEL_DEBUG
	case string(_LEVEL_INFO):
		return LEVEL_INFO
	case string(_LEVEL_WARNING):
		fallthrough
	case "WARNING":
		return LEVEL_WARNING
	case string(_LEVEL_ERROR):
		return LEVEL_ERROR
	case string(_LEVEL_FATAL):
		return LEVEL_FATAL
	default:
		Errorf("could not parse logLevel from string %s", levelRaw)
		return LEVEL_INFO
	}
}
