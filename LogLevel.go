package ulog

type LogLevel int

const (
	LEVEL_TRACE LogLevel = iota
	LEVEL_DEBUG
	LEVEL_INFO
	LEVEL_WARNING
	LEVEL_ERROR
	LEVEL_FATAL
)

var logLevel LogLevel = LEVEL_INFO

func GetLogLevel() LogLevel {
	return logLevel
}

func SetLogLevel(_level LogLevel) {
	logLevel = _level
}

func SetLogLevelFromString(levelRaw string) {
	SetLogLevel(LogLevelFromString(levelRaw))
}
