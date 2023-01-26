module github.com/dunv/ulog/v2

go 1.19

require go.uber.org/zap v1.24.0-traceLevel.1

require (
	go.uber.org/atomic v1.7.0 // indirect
	go.uber.org/multierr v1.6.0 // indirect
)

// a fork of zap which has traceLevel support
replace go.uber.org/zap v1.24.0-traceLevel.1 => github.com/dunv/zap v1.24.0-traceLevel.1
