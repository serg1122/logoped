package logoped

import "time"

// Report
type IReport interface {
	GetHeader() IHeader
	GetMessage() []byte
	GetPostMessage() []byte
}

type DefaultReport struct {
	IReport

	header      IHeader
	message     []byte
	postmessage []byte
}

func CreateDefaultReport(header IHeader, message []byte, postmessage []byte) *DefaultReport {

	return &DefaultReport{
		header:      header,
		message:     message,
		postmessage: postmessage,
	}
}

func (report DefaultReport) GetHeader() IHeader {

	return report.header
}

func (report DefaultReport) GetMessage() []byte {

	return report.message
}

func (report DefaultReport) GetPostMessage() []byte {

	return report.postmessage
}

// Header
type IHeader interface {
	GetTimepoint() time.Time
	GetLogLevel() ILogLevel
}

type Header struct {
	timepoint time.Time
	loglevel  ILogLevel
}

func CreateHeader(timepoint time.Time, loglevel ILogLevel) *Header {

	return &Header{
		timepoint: timepoint,
		loglevel:  loglevel,
	}
}

func (header Header) GetTimepoint() time.Time {

	return header.timepoint
}

func (header Header) GetLogLevel() ILogLevel {

	return header.loglevel
}
