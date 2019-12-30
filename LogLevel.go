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

// Set default logLevel
var logLevel LogLevel = LEVEL_INFO

// Get currently set logLevel
func GetLogLevel() LogLevel {
	return logLevel
}

// Set logLevel
func SetLogLevel(_level LogLevel) {
	logLevel = _level
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
	_LEVEL_TRACE   logLevelString = "TRACE"
	_LEVEL_DEBUG   logLevelString = "DEBUG"
	_LEVEL_INFO    logLevelString = "INFO "
	_LEVEL_WARNING logLevelString = "WARN "
	_LEVEL_ERROR   logLevelString = "ERROR"
	_LEVEL_FATAL   logLevelString = "FATAL"
)
