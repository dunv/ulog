package ulog

var globalDebug bool = false

// Enable debugging skipFunctions and ReplaceFunctions
func SetDebug() {
	globalDebug = true
}

// Disable debugging skipFunctions and ReplaceFunctions
func UnsetDebug() {
	globalDebug = false
}
