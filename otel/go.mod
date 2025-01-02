module github.com/dunv/ulog/otel

go 1.23.2

require (
	go.opentelemetry.io/otel/log v0.9.0
	go.uber.org/zap v1.27.0
)

require (
	github.com/go-logr/logr v1.4.2 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	go.opentelemetry.io/auto/sdk v1.1.0 // indirect
	go.opentelemetry.io/otel v1.33.0 // indirect
	go.opentelemetry.io/otel/metric v1.33.0 // indirect
	go.opentelemetry.io/otel/trace v1.33.0 // indirect
	go.uber.org/multierr v1.10.0 // indirect
)

replace go.uber.org/zap v1.27.0 => github.com/dunv/zap v1.27.0-traceLevel.2
