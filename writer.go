package ulog

type DebugWriter struct{}

func (DebugWriter) Write(p []byte) (int, error) {
	SkipOneS().Debug(string(p))
	return len(p), nil
}

type InfoWriter struct{}

func (InfoWriter) Write(p []byte) (int, error) {
	SkipOneS().Info(string(p))
	return len(p), nil
}

type WarnWriter struct{}

func (WarnWriter) Write(p []byte) (int, error) {
	SkipOneS().Warn(string(p))
	return len(p), nil
}

type ErrorWriter struct{}

func (ErrorWriter) Write(p []byte) (int, error) {
	SkipOneS().Error(string(p))
	return len(p), nil
}

type FatalWriter struct{}

func (FatalWriter) Write(p []byte) (int, error) {
	SkipOneS().Fatal(string(p))
	return len(p), nil
}

type PanicWriter struct{}

func (PanicWriter) Write(p []byte) (int, error) {
	SkipOneS().Panic(string(p))
	return len(p), nil
}
