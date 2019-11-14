package logoped

import (
	"sync"
	"time"
)

type IDispatcher interface {
	CreateReport(ILogLevel, []byte) IReport
	Dispatch(IReport)
}

type Dispatcher struct {
	acceptors []IAcceptor
}

func CreateDispatcher(acceptors []IAcceptor) *Dispatcher {

	return &Dispatcher{
		acceptors: acceptors,
	}
}

func (d Dispatcher) CreateReport(loglevel ILogLevel, bytes []byte) IReport {

	return CreateReport(CreateHeader(time.Now().UTC(), loglevel), bytes)
}

func (d Dispatcher) Dispatch(report IReport) {

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
