package ulog

type logLevelString string

const (
	_LEVEL_TRACE   logLevelString = "TRACE"
	_LEVEL_DEBUG   logLevelString = "DEBUG"
	_LEVEL_INFO    logLevelString = "INFO "
	_LEVEL_WARNING logLevelString = "WARN "
	_LEVEL_ERROR   logLevelString = "ERROR"
	_LEVEL_FATAL   logLevelString = "FATAL"
)
