module github.com/dunv/ulog/v2

go 1.22

require go.uber.org/zap v1.27.0

require go.uber.org/multierr v1.11.0 // indirect

// a fork of zap which has traceLevel support
replace go.uber.org/zap v1.27.0 => ../../../Git/zap

// replace go.uber.org/zap v1.24.0-traceLevel.1 => github.com/dunv/zap v1.24.0-traceLevel.1
