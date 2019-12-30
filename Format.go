package ulog

import "text/template"

var origFmtString = "{{ .Time }} | {{ .Level }} | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n"
var lineTemplate *template.Template = template.Must(template.New("lineTemplate").Parse(origFmtString))
var origTsFormat string = "2006-01-02 15:04:05.000"

var fmtString = origFmtString
var tsFormat = origTsFormat

// Use original logging format
func ResetFormat() {
	fmtString = origFmtString
	lineTemplate = template.Must(template.New("lineTemplate").Parse(origFmtString))
	tsFormat = origTsFormat
}

// Use custom logging format (using text/template)
// the original is "{{ .Time }} | {{ .Level }} | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n"
// Available vars are
// - Time
// - Level
// - Package
// - File (not always available)
// - Line (not always available)
// - Function (not always available)
// - Message
func SetFormatString(_fmtString string) {
	fmtString = _fmtString
	lineTemplate = template.Must(template.New("lineTemplate").Parse(fmtString))
}

// Set a different format for the timestamp
// the original is "2006-01-02 15:04:05.000"
func SetTimestampFormat(_tsFormat string) {
	tsFormat = _tsFormat
}
