package ulog

type LogEntry struct {
	Time     string
	Level    logLevelString
	Package  string
	File     string
	Line     int
	Function string
	Message  string
}
