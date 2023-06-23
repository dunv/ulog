package v2

type TraceWriter struct{}

func (TraceWriter) Write(p []byte) (int, error) {
	skipOneSugaredLogger.Trace(string(p))
	return len(p), nil
}

type DebugWriter struct{}

func (DebugWriter) Write(p []byte) (int, error) {
	skipOneSugaredLogger.Debug(string(p))
	return len(p), nil
}

type InfoWriter struct{}

func (InfoWriter) Write(p []byte) (int, error) {
	skipOneSugaredLogger.Info(string(p))
	return len(p), nil
}

type WarnWriter struct{}

func (WarnWriter) Write(p []byte) (int, error) {
	skipOneSugaredLogger.Warn(string(p))
	return len(p), nil
}

type ErrorWriter struct{}

func (ErrorWriter) Write(p []byte) (int, error) {
	skipOneSugaredLogger.Error(string(p))
	return len(p), nil
}

type FatalWriter struct{}

func (FatalWriter) Write(p []byte) (int, error) {
	skipOneSugaredLogger.Fatal(string(p))
	return len(p), nil
}

type PanicWriter struct{}

func (PanicWriter) Write(p []byte) (int, error) {
	skipOneSugaredLogger.Panic(string(p))
	return len(p), nil
}
