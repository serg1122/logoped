package logoped

import (
	"io"
)

type IAcceptor interface {
	GetFormatter() IFormatter
	GetLogLevel() ILogLevel
	GetWriters() []io.Writer
}

type Acceptor struct {
	formatter IFormatter
	logLevel  ILogLevel
	writers   []io.Writer
}

func CreateAcceptor(formatter IFormatter, logLevel ILogLevel, writers []io.Writer) *Acceptor {

	return &Acceptor{
		formatter: formatter,
		logLevel:  logLevel,
		writers:   writers,
	}
}

func (acceptor Acceptor) GetLogLevel() ILogLevel {

	return acceptor.logLevel
}

func (acceptor Acceptor) GetFormatter() IFormatter {

	return acceptor.formatter
}

func (acceptor Acceptor) GetWriters() []io.Writer {

	return acceptor.writers
}
