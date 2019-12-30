package ulog

var debug bool = false

// Enable debugging skipFunctions and ReplaceFunctions
func SetDebug() {
	debug = true
}

// Disable debugging skipFunctions and ReplaceFunctions
func UnsetDebug() {
	debug = false
}
