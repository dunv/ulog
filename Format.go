package ulog

import (
	"text/template"
)

var origFmtString = "{{ .Time }} | {{ .Level }} | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n"
var globalLineTemplate *template.Template = template.Must(template.New("lineTemplate").Parse(origFmtString))
var origTsFormat string = "2006-01-02 15:04:05.000"

var globalTsFormat = origTsFormat

// Use original logging format
func ResetFormat() {
	globalLineTemplate = template.Must(template.New("lineTemplate").Parse(origFmtString))
	globalTsFormat = origTsFormat
}

// Use custom logging format (using text/template)
// The original is
//  "{{ .Time }} | {{ .Level }} | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n"
// Available vars are
//  - Time
//  - Level
//  - Package
//  - File (not always available)
//  - Line (not always available)
//  - Function (not always available)
//  - Message
func SetFormatString(_fmtString string) {
	globalLineTemplate = template.Must(template.New("lineTemplate").Parse(_fmtString))
}

// Set a different format for the timestamp
// the original is "2006-01-02 15:04:05.000"
func SetTimestampFormat(_tsFormat string) {
	globalTsFormat = _tsFormat
}
