[![Build Status](https://travis-ci.org/dunv/ulog.svg?branch=master)](https://travis-ci.org/dunv/ulog)
[![GoDoc](https://godoc.org/github.com/dunv/ulog?status.svg)](https://godoc.org/github.com/dunv/ulog)
[![codecov](https://codecov.io/gh/dunv/ulog/branch/master/graph/badge.svg)](https://codecov.io/gh/dunv/ulog)

# ulog

A simple logging library which comes with an extension to log to influxDB (https://github.com/dunv/ulog2influx)

It comes with a default template which prints log-lines for example like this
```
2019-12-30 20:42:20.372 | ERROR | github.com/dunv/uauth/handlers.glob. CheckLoginHandler.go:59 (func1) | renderError [path: /api/checkLogin] token is expired by 141h11m24s
2019-12-30 20:42:20.372 | INFO  | uhttp.Handle | Registered http GET /api/monthReport
2019-12-30 20:42:20.372 | INFO  | main main.go:105 (main) | Serving at 0.0.0.0:8080
```

### Setup
```
ulog.SetLogLevelFromString(cfg.LogLevel)
```
- Additionally one can set up functions to skip (see doc)
- Additionally one can set up functions to replace (see doc)

### Usage examples
```
ulog.Trace("helloWorld")
ulog.Tracef("helloWorld %s", "today")
ulog.Debug("helloWorld")
ulog.Infof("helloWorld %s", "today")
ulog.Errorf("helloWorld %s", "today")
ulog.Fatalf("helloWorld %s", "today")
ulog.Panicf("helloWorld %s", "today")
```

### Readable console fields

Enable key/value console fields explicitly:

```go
ulog.Configure(
    ulog.WithStripAdditionalFields(false),
    ulog.WithConsoleKeyValues(true),
)
```

For example: `Failed to connect | destination=inventory attempt=3 error="connection refused"`.
Keys and `=` are dark grey; values use the terminal's default foreground. Existing
message/level colors remain. `WithColors(false)` disables all console ANSI colors.
Strings are quoted when empty or containing whitespace or special characters;
JSON escaping keeps field newlines and terminal control sequences on one line.
Nested objects and arrays remain compact JSON values. Field order, duplicate keys,
and numeric precision are preserved, including persistent `.With(...)` fields.
Additional structured cores are unchanged.

The default remains the existing JSON suffix when fields are enabled. Keep that
format for consumers that parse console JSON. `WithStripAdditionalFields(true)`
still suppresses fields in either format. The obsolete dummy-thread column has
been removed.

Compare encoder costs (including allocations and colored output) with:

```sh
go test -run '^$' -bench BenchmarkConsoleFields -benchmem -count=5 ./...
```

The benchmark covers no fields, error context, persistent fields, and nested
payloads, comparing JSON, key/value, and stripped-field modes. It measures encoding,
not terminal I/O. The key/value formatter scans Zap's JSON buffer directly without
allocating a map or decoding field values.

### Wrapping configured cores

`WithCoreWrapper(func(zapcore.Core) zapcore.Core)` wraps the combined console and
additional cores once during `Configure`, after `WithBaseFields` has attached its
fields. Both Zap's global logger and ulog's cached/fluent loggers use the wrapped
core. This allows field guards or hooks without configuring a second console sink.

The wrapper must return a non-nil core. Passing `nil` disables wrapping; the last
`WithCoreWrapper` option wins. Like other options, it is reset by a subsequent
`Configure` unless supplied again. Existing configurations without this option are
unchanged.