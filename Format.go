package ulog

import "text/template"

var fmtString = "{{ .Time }} | {{ .Level }} | n/a | {{ .Package }}{{ if .File }} {{ .File }}{{ end }}{{ if .Line }}:{{ .Line }}{{ end }}{{ if .Function }} ({{ .Function }}){{ end }} | {{ .Message }}\n"
var lineTemplate *template.Template = template.Must(template.New("lineTemplate").Parse(fmtString))
var tsFormat string = "2006-01-02 15:04:05.000"

func SetFormatString(_fmtString string) {
	fmtString = _fmtString
	lineTemplate = template.Must(template.New("lineTemplate").Parse(fmtString))
}

func SetTimestampFormat(_tsFormat string) {
	tsFormat = _tsFormat
}
