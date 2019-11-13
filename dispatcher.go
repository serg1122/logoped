package logoped

import (
	"sync"
	"time"
)

type IDispatcher interface {
	CreateReport(ILogLevel, []byte, []byte) IReport
	Dispatch(IReport)
}

type DefaultDispatcher struct {
	acceptors []IAcceptor
}

func CreateDefaultDispatcher(acceptors []IAcceptor) *DefaultDispatcher {

	return &DefaultDispatcher{
		acceptors: acceptors,
	}
}

func (d DefaultDispatcher) CreateReport(loglevel ILogLevel, bytes []byte, postmessage []byte) IReport {

	return CreateDefaultReport(CreateHeader(time.Now().UTC(), loglevel), bytes, postmessage)
}

func (d DefaultDispatcher) Dispatch(report IReport) {

	var wg sync.WaitGroup

	for _, acceptor := range d.acceptors {

		if report.GetHeader().GetLogLevel().Weight() >= acceptor.GetLogLevel().Weight() {

			bytes := acceptor.GetFormatter().Format(report)

			for _, writer := range acceptor.GetWriters() {

				currentWriter := writer

				var process = func() {

					defer wg.Done()

					currentWriter.Write(bytes)
				}

				wg.Add(1)
				go process()
			}
		}
	}

	wg.Wait()
}
