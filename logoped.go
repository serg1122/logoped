package logoped

var LogLevelInfo = CreateLogLevel("INFO", 100)
var LogLevelWarn = CreateLogLevel("WARN", 200)
var LogLevelError = CreateLogLevel("ERROR", 300)
var LogLevelFatal = CreateLogLevel("FATAL", 400)

type Logoped struct {
	logLevel ILogLevel
}

func CreateLogoped(logLevel ILogLevel) *Logoped {

	return &Logoped{
		logLevel: logLevel,
	}
}

// func (l Logoped) CreateMessage(bytes []byte) *Report {
// 	//...
// }

// func (l Logoped) Info(bytes []byte) {

// 	if l.logLevel.Weight() < LogLevelInfo.Weight() {
// 		return
// 	}

// 	l.CreateMessage(bytes)
// }
