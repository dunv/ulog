[![Build Status](https://travis-ci.org/dunv/ulog.svg?branch=master)](https://travis-ci.org/dunv/ulog)
[![GoDoc](https://godoc.org/github.com/dunv/ulog?status.svg)](https://godoc.org/github.com/dunv/ulog)

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
ulog.Tracef("helloWorld %s", "today)
ulog.Debug("helloWorld")
ulog.Infof("helloWorld %s", "today)
ulog.Errorf("helloWorld %s", "today)
ulog.Fatalf("helloWorld %s", "today)
ulog.Panicf("helloWorld %s", "today)
```