package logoped

import "runtime/debug"

var LogLevelInfo = CreateLogLevel("INFO", 100)
var LogLevelWarn = CreateLogLevel("WARN", 200)
var LogLevelError = CreateLogLevel("ERROR", 300)
var LogLevelFatal = CreateLogLevel("FATAL", 400)

type Logoped struct {
	dispatcher IDispatcher
	logLevel   ILogLevel
}

func CreateLogoped(dispatcher IDispatcher, logLevel ILogLevel) *Logoped {

	return &Logoped{
		dispatcher: dispatcher,
		logLevel:   logLevel,
	}
}

func (l Logoped) Log(report IReport) {

	l.dispatcher.Dispatch(report)
}

func (l Logoped) Info(bytes []byte) {

	if l.logLevel.Weight() < LogLevelInfo.Weight() {
		return
	}

	l.dispatcher.Dispatch(l.dispatcher.CreateReport(LogLevelInfo, bytes, []byte{}))
}

func (l Logoped) Warn(bytes []byte) {

	if l.logLevel.Weight() > LogLevelWarn.Weight() {
		return
	}

	l.dispatcher.Dispatch(l.dispatcher.CreateReport(LogLevelWarn, bytes, []byte{}))
}

func (l Logoped) Error(bytes []byte) {

	if l.logLevel.Weight() > LogLevelError.Weight() {
		return
	}

	l.dispatcher.Dispatch(l.dispatcher.CreateReport(LogLevelError, bytes, debug.Stack()))
}

func (l Logoped) Fatal(bytes []byte) {

	if l.logLevel.Weight() < LogLevelFatal.Weight() {
		return
	}

	l.dispatcher.Dispatch(l.dispatcher.CreateReport(LogLevelFatal, bytes, debug.Stack()))
}
