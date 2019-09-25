package ulog

import (
	"io"
	"os"
)

var writer io.Writer = os.Stdout

func SetWriter(_writer io.Writer) {
	writer = _writer
	logger.SetOutput(_writer)
}
