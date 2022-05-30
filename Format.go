package ulog

import (
	"strings"
	"text/template"
)

var colors = []string{
	"{{ if eq .Level \"TRACE\" }}\033[90m{{ end }}",
	"{{ if eq .Level \"DEBUG\" }}\033[34m{{ end }}",
	"{{ if eq .Level \"INFO \" }}\033[97m{{ end }}",
	"{{ if eq .Level \"WARN \" }}\033[95m{{ end }}",
	"{{ if eq .Level \"ERROR\" }}\033[91m{{ end }}",
	"{{ if eq .Level \"FATAL\" }}\033[93m{{ end }}",
}

var coloredFmtString = strings.Join(append(colors, monochromeFmtString, "\033[0m"), "")
var monochromeFmtString = "{{ .Time }} | {{ .Level }} | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}"
var tsFormatString string = "2006-01-02 15:04:05.000"

var globalLineTemplate *template.Template = template.Must(template.New("lineTemplate").Parse(monochromeFmtString))
var globalTsFormat = tsFormatString

// Use a format which colors lines according to logLevel
func EnableColors() {
	globalLineTemplate = template.Must(template.New("lineTemplate").Parse(coloredFmtString))
}

// Use original logging format
func ResetFormat() {
	globalLineTemplate = template.Must(template.New("lineTemplate").Parse(monochromeFmtString))
	globalTsFormat = tsFormatString
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
