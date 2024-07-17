package v2

import (
	"bytes"
	"fmt"
	"io"
	"regexp"
	"testing"

	"github.com/stretchr/testify/require"
	"go.uber.org/zap/zapcore"
)

func lines(t *testing.T, b *bytes.Buffer) [][]byte {
	out, err := io.ReadAll(b)
	require.NoError(t, err)
	parts := bytes.Split(out, []byte("\n"))
	return parts[:len(parts)-1]
}

func requireMsgEquals(t *testing.T, line []byte, level zapcore.Level, msg string) {
	colorStart := regexp.QuoteMeta(string(levelToColorStart[level]))
	colorEnd := regexp.QuoteMeta(string(levelToColorEnd[level]))
	msg = regexp.QuoteMeta(msg)

	lvlString := levelString[level]
	regexString := fmt.Sprintf(`(?m)^%s\d\d\d\d-\d\d-\d\dT\d\d:\d\d:\d\d\+\d\d:\d\d \| %s \| [\w\/]+\.go:\d+ \| %s%s$`, colorStart, lvlString, msg, colorEnd)
	re := regexp.MustCompile(regexString)
	require.True(t, re.Match(line), `did not log the correct output actual:"%s" regex:"%s"`, string(line), regexString)
}
