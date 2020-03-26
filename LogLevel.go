package ulog

import "strings"

// LogLevel for the application
type LogLevel int

const (
	LEVEL_TRACE LogLevel = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARNING
	LEVEL_ERROR
	LEVEL_FATAL
)

func (l LogLevel) String() string {
	switch l {
	case LEVEL_TRACE:
		return string(_LEVEL_TRACE)
	case LEVEL_DEBUG:
		return string(_LEVEL_DEBUG)
	case LEVEL_INFO:
		return string(_LEVEL_INFO)
	case LEVEL_WARNING:
		return string(_LEVEL_WARNING)
	case LEVEL_ERROR:
		return string(_LEVEL_ERROR)
	case LEVEL_FATAL:
		return string(_LEVEL_FATAL)
	default:
		return string(_LEVEL_UNDEFINED)
	}
}

// Set default logLevel
var globalLogLevel LogLevel = LEVEL_INFO

// Get currently set logLevel
func GetLogLevel() LogLevel {
	return globalLogLevel
}

// Set logLevel
func SetLogLevel(_level LogLevel) {
	globalLogLevel = _level
}

// Set logLevel from string
// takes "trace", "debug", "info", "information", "warn", "warning", "error" and "fatal" as arguments. Case does not matter
func SetLogLevelFromString(levelRaw string) {
	SetLogLevel(logLevelFromString(levelRaw))
}

// Helper function to parse a log-level received as string into the correct level
func logLevelFromString(levelRaw string) LogLevel {
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

// Internal type name (this way all log-level identifiers have the same width)
type logLevelString string

const (
	_LEVEL_TRACE     logLevelString = "TRACE"
	_LEVEL_DEBUG     logLevelString = "DEBUG"
	_LEVEL_INFO      logLevelString = "INFO "
	_LEVEL_WARNING   logLevelString = "WARN "
	_LEVEL_ERROR     logLevelString = "ERROR"
	_LEVEL_FATAL     logLevelString = "FATAL"
	_LEVEL_UNDEFINED logLevelString = "UNDEFINED"
)
